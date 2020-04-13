package controllers

import (
	"errors"
	"github.com/gorilla/context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"webserverREST/internal/constants"
	"webserverREST/internal/model"
	"webserverREST/internal/tools"
)

const ID string = "id"

type TestData struct {
	name              string
	userID            *string
	userName          string
	userLastname      string
	userAge           int
	userBirthdate     time.Time
	repositoryReturns bool
	expectStatusCode  int
}

var testsPut = []TestData{
	{
		name:              "valid user model",
		userID:            tools.GenUUID(),
		userName:          "James",
		userLastname:      "Young",
		userBirthdate:     time.Now(),
		repositoryReturns: true,
		expectStatusCode:  200,
	},
	{
		name:              "invalid user model",
		userID:            nil,
		userName:          "",
		userLastname:      "Young",
		userBirthdate:     time.Now(),
		repositoryReturns: true,
		expectStatusCode:  400,
	},
	{
		name:              "invalid repository response",
		userID:            tools.GenUUID(),
		userName:          "James",
		userLastname:      "Young",
		userBirthdate:     time.Now(),
		repositoryReturns: false,
		expectStatusCode:  500,
	},
}

var testsGet = []TestData{
	{
		name:              "invalid repository response",
		userID:            tools.GenUUID(),
		userName:          "James",
		userLastname:      "Young",
		userBirthdate:     time.Now(),
		repositoryReturns: true,
		expectStatusCode:  200,
	},
	{
		name:              "valid repository response",
		userID:            tools.GenUUID(),
		userName:          "James",
		userLastname:      "Young",
		userBirthdate:     time.Now(),
		repositoryReturns: false,
		expectStatusCode:  500,
	},
}

func TestCreateController(t *testing.T) {
	for _, test := range testsPut {
		t.Run(test.name, func(t *testing.T) {
			req, repo := setupTests(test)
			controller := NewUser(repo)
			w := httptest.NewRecorder()
			controller.Create(w, req)
			assert.Equal(t, test.expectStatusCode, w.Code)
		})
	}
}

func TestEditController(t *testing.T) {
	for _, test := range testsPut {
		t.Run(test.name, func(t *testing.T) {
			req, repo := setupTests(test)
			controller := NewUser(repo)
			w := httptest.NewRecorder()
			controller.Edit(w, req)
			assert.Equal(t, test.expectStatusCode, w.Code)
		})
	}
}

func TestGetController(t *testing.T) {
	for _, test := range testsGet {
		t.Run(test.name, func(t *testing.T) {
			req, repo := setupTests(test)
			controller := NewUser(repo)
			w := httptest.NewRecorder()
			controller.Get(w, req)
			assert.Equal(t, test.expectStatusCode, w.Code)
		})
	}
}

func TestDeleteController(t *testing.T) {
	for _, test := range testsGet {
		t.Run(test.name, func(t *testing.T) {
			req, repo := setupTests(test)
			controller := NewUser(repo)
			w := httptest.NewRecorder()
			controller.Delete(w, req)
			assert.Equal(t, test.expectStatusCode, w.Code)
		})
	}
}

func setupTests(data TestData) (*http.Request, *MockedRepository) {
	return prepareContext(data), prepareRepository(data)
}

func prepareContext(data TestData) *http.Request {
	var (
		userName     *string
		userLastname *string
	)
	req, _ := http.NewRequest("GET", "/users", nil)
	if len(data.userName) == 0 {
		userName = nil
	} else {
		userName = &data.userName
	}
	if len(data.userLastname) == 0 {
		userLastname = nil
	} else {
		userLastname = &data.userLastname
	}
	model := model.UserModel{
		Id:        data.userID,
		Name:      userName,
		Lastname:  userLastname,
		Age:       tools.Age(data.userBirthdate),
		Birthdate: &data.userBirthdate,
	}
	context.Set(req, constants.Body, &model)
	context.Set(req, ID, data.userID)
	return req
}

func prepareRepository(data TestData) *MockedRepository {
	re := new(MockedRepository)
	if data.repositoryReturns {
		re.On("Put", mock.Anything).Return(nil)
		re.On("Delete", mock.Anything).Return(nil)
		re.On("Get").Return([]model.UserModel{}, nil)
	} else {
		re.On("Put", mock.Anything).Return(errors.New(""))
		re.On("Delete", mock.Anything).Return(errors.New(""))
		re.On("Get").Return([]model.UserModel{}, errors.New(""))
	}
	return re
}

type MockedRepository struct {
	mock.Mock
}

func (m *MockedRepository) Put(u *model.UserModel) error {
	args := m.Called(u)
	return args.Error(0)
}

func (m *MockedRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockedRepository) Get() ([]model.UserModel, error) {
	args := m.Called()
	return []model.UserModel{}, args.Error(1)
}
