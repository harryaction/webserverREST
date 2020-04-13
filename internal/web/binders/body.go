package binders

import (
	"encoding/json"
	"github.com/gorilla/context"
	"io/ioutil"
	"log"
	"net/http"
	"webserverREST/internal/constants"
	"webserverREST/internal/model"
)

func BodyParse(r *http.Request) {
	data := &model.UserModel{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
	} else {
		err = json.Unmarshal([]byte(body), data)
		if err != nil {
			log.Printf("Error parsing body: %v", err)
		} else {
			context.Set(r, constants.Body, data)
		}
	}
}
