package test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"
	"time"
	"webserverREST/internal"
	"webserverREST/internal/model"
)

type Suite struct {
	api internal.Api
	suite.Suite
	DB     *sqlx.DB
	mock   sqlmock.Sqlmock
	person model.UserModel
}

func (s *Suite) SetupSuite() {
	var (
		dbm *sql.DB
		err error
	)
	dbm, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)
	s.DB = sqlx.NewDb(dbm, "sqlmock")
	require.NoError(s.T(), err)
	s.api = internal.Api{DB: s.DB}
	s.api.Setup()
	require.NoError(s.T(), err)
}

func (s *Suite) TestEmptyTable() {
	req, _ := http.NewRequest("GET", "/users", nil)
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT uuid id, name, lastname, birthdate, 0 age FROM public.api_users`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "lastname", "birthdate", "age"}))
	response := executeRequest(s.api, req)
	assert.Equal(s.T(), http.StatusOK, response.Code)
	assert.Equal(s.T(), "[]", strings.TrimSuffix(response.Body.String(), "\n"))
}

func (s *Suite) TestCreate() {
	var jsonStr = []byte(`{"name":"Trevor", "lastname": "Young", "birthdate": "2005-05-05T00:00:00Z"}`)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStr))
	s.mock.ExpectExec(regexp.QuoteMeta(
		`INSERT INTO public.api_users(uuid, 
                             name, 
                             lastname, 
                             birthdate) 
                             VALUES($1, $2, $3, $4::timestamptz) ON CONFLICT(uuid) DO UPDATE SET name = $2, lastname = $3, birthdate = $4`)).WithArgs(sqlmock.AnyArg(), "Trevor", "Young", sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))
	response := executeRequest(s.api, req)
	assert.Equal(s.T(), http.StatusOK, response.Code)
	err := json.Unmarshal([]byte(response.Body.String()), &model.UserModel{})
	require.NoError(s.T(), err)
}

func (s *Suite) TestPut() {
	var jsonStr = []byte(`{"id": "ebdded24-d979-485d-b1cb-4ae179da787c", "name": "Trevor", "lastname": "Young", "birthdate": "2005-05-05T00:00:00Z"}`)
	req, _ := http.NewRequest("PUT", "/users", bytes.NewBuffer(jsonStr))
	s.mock.ExpectExec(regexp.QuoteMeta(
		`INSERT INTO public.api_users(uuid, 
                             name, 
                             lastname, 
                             birthdate) 
                             VALUES($1, $2, $3, $4::timestamptz) ON CONFLICT(uuid) DO UPDATE SET name = $2, lastname = $3, birthdate = $4`)).WithArgs(sqlmock.AnyArg(), "Trevor", "Young", sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(0, 1))
	response := executeRequest(s.api, req)
	assert.Equal(s.T(), http.StatusOK, response.Code)
	responseData := &model.UserModel{}
	err := json.Unmarshal([]byte(response.Body.String()), responseData)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), "ebdded24-d979-485d-b1cb-4ae179da787c", *responseData.Id)
}

func (s *Suite) TestGet() {
	var (
		id       = "ebdded24-d979-485d-b1cb-4ae179da787c"
		name     = "Trevor"
		lastname = "Young"
		str      = "2005-05-05T00:00:00Z"
	)
	layout := "2006-01-02T15:04:05.000Z"
	t, _ := time.Parse(layout, str)
	req, _ := http.NewRequest("GET", "/users", nil)
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT uuid id, name, lastname, birthdate, 0 age FROM public.api_users`)).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "lastname", "birthdate", "age"}).AddRow(id, name, lastname, t, 0))
	response := executeRequest(s.api, req)
	assert.Equal(s.T(), http.StatusOK, response.Code)
	responseData := &[]model.UserModel{}
	err := json.Unmarshal([]byte(response.Body.String()), responseData)
	require.NoError(s.T(), err)
}

func (s *Suite) TestDelete() {
	req, _ := http.NewRequest("DELETE", "/users/ebdded24-d979-485d-b1cb-4ae179da787c", nil)
	s.mock.ExpectExec(regexp.QuoteMeta(
		`DELETE FROM public.api_users WHERE uuid = $1`)).WithArgs(sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(0, 1))
	response := executeRequest(s.api, req)
	assert.Equal(s.T(), http.StatusOK, response.Code)
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func executeRequest(app internal.Api, req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, req)
	return rr
}

func TestApiTestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
