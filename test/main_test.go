package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"webserverREST/internal/model"
	"webserverREST/internal/web/controllers"
)

var jsonStr = []byte(`{"Name":"Jerry", "Lastname": "King", "Birthdate": "1985-09-04T00:00:00Z"}`)
var url = "http://some.host/user/"

func TestServerCanReturnData(t *testing.T) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	controllers.CreateUser(res, req)
	act := res.Body.String()
	data := &model.UserModel{}
	err = json.Unmarshal([]byte(act), data)
	if err != nil {
		t.Fatalf("Expected correct JSON in response, responded with %s", act)
	}
	if data.Name != "Somename" || data.Lastname != "Surname" {
		t.Fatalf("Expected Name Surname, responded with %s %s", data.Name, data.Lastname)
	}
	fmt.Println(data.Id)
	fmt.Println(data.Age)
}
