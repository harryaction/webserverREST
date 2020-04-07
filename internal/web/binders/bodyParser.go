package binders

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"webserverREST/internal/model"
	"webserverREST/internal/tools"
)

func Parse(r *http.Request, create bool, update bool) (*model.UserModel, error) {
	data := &model.UserModel{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		return data, fmt.Errorf("error reading body: %v", err)
	} else {
		err = json.Unmarshal([]byte(body), data)
		if err != nil {
			log.Printf("Error parsing body: %v", err)
			return data, fmt.Errorf("error parsing body: %v", err)
		} else {
			if create {
				data.Id = tools.GenUUID()
			}
			if create || update {
				data.Age = tools.Age(data.Birthdate)
			}
			return data, nil
		}
	}
}
