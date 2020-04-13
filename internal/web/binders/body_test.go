package binders

import (
	"bytes"
	"github.com/gorilla/context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"webserverREST/internal/constants"
	"webserverREST/internal/model"
)

type TestCase struct {
	name       string
	jsonStr    string
	expectName string
}

var tests = []TestCase{
	{
		name:       "broken json in request",
		jsonStr:    "{\"last\": \"Young\", \"birth\": \"2005-05-05T00:00:00Z\"}",
		expectName: "",
	},
	{
		name:       "valid json in request",
		jsonStr:    "{\"id\": \"ebdded24-d979-485d-b1cb-4ae179da787c\", \"name\": \"Trevor\", \"lastname\": \"Young\", \"birthdate\": \"2005-05-05T00:00:00Z\"}",
		expectName: "Trevor",
	},
	{
		name:       "empty json in request",
		jsonStr:    "",
		expectName: "",
	},
}

func TestBodyBinder(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var (
				expect *string
				actual *string
			)
			req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer([]byte(test.jsonStr)))
			BodyParse(req)

			if len(test.expectName) == 0 {
				expect = nil
			} else {
				expect = &test.expectName
			}

			if response := context.Get(req, constants.Body); response != nil {
				actual = response.(*model.UserModel).Name
			} else {
				actual = nil
			}
			assert.Equal(t, expect, actual)
		})
	}
}
