package repositories

import (
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"regexp"
	"testing"
	"time"
	"webserverREST/internal/model"
	"webserverREST/internal/tools"
)

type TestData struct {
	userID        *string
	userName      string
	userLastname  string
	userBirthdate time.Time
}

var tests = []TestData{
	{
		userID:        tools.GenUUID(),
		userName:      "James",
		userLastname:  "Young",
		userBirthdate: time.Now(),
	},
	{
		userID:        nil,
		userName:      "",
		userLastname:  "Young",
		userBirthdate: time.Now(),
	},
}

type Suite struct {
	repository User
	suite.Suite
	DB     *sqlx.DB
	mock   sqlmock.Sqlmock
	person *model.UserModel
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
	s.repository = NewUser(s.DB)
	require.NoError(s.T(), err)
	s.person = &model.UserModel{
		Id:        tools.GenUUID(),
		Name:      nil,
		Lastname:  nil,
		Age:       0,
		Birthdate: nil,
	}
}

func (s *Suite) TestRepositoryPut(t *testing.T) {
	s.mock.ExpectExec(regexp.QuoteMeta(
		`INSERT INTO public.api_users(uuid, 
                             name, 
                             lastname, 
                             birthdate) 
                             VALUES($1, $2, $3, $4::timestamptz) ON CONFLICT(uuid) DO UPDATE SET name = $2, lastname = $3, birthdate = $4`)).WithArgs(sqlmock.AnyArg(), nil, sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnError(errors.New("err"))
	for _, testCase := range s.person {
		err := s.repository.Put(&testCase)
	}
}

func TestApiTestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
