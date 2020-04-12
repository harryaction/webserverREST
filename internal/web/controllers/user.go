package controllers

import (
	"encoding/json"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"net/http"
	"webserverREST/internal/model"
	"webserverREST/internal/repositories"
	"webserverREST/internal/tools"
	"webserverREST/internal/web/binders"
)

type User interface {
	Create(w http.ResponseWriter, req *http.Request)
	Edit(w http.ResponseWriter, req *http.Request)
	Get(w http.ResponseWriter, req *http.Request)
	Delete(w http.ResponseWriter, req *http.Request)
}

type user struct {
	repo repositories.User
}

func NewUser(repo repositories.User) User {
	return &user{repo: repo}
}

func (u user) Create(w http.ResponseWriter, req *http.Request) {
	if response := context.Get(req, binders.Body); response != nil {
		response := response.(*model.UserModel)
		if (response.Name != nil) && (response.Lastname != nil) && (response.Birthdate != nil) {
			response.Id = tools.GenUUID()
			response.Age = tools.Age(*response.Birthdate)
			err := u.repo.Put(response)
			if err != nil {
				http.Error(w, "DB error", http.StatusInternalServerError)
			} else {
				w.WriteHeader(200)
				json.NewEncoder(w).Encode(response)
			}
		} else {
			http.Error(w, "JSON parsing error", http.StatusBadRequest)
		}
	} else {
		http.Error(w, "JSON parsing error", http.StatusBadRequest)
	}
}

func (u user) Delete(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := vars["id"]
	err := u.repo.Delete(id)
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
	} else {
		w.WriteHeader(200)
	}
}

func (u user) Edit(w http.ResponseWriter, req *http.Request) {
	if response := context.Get(req, binders.Body); response != nil {
		response := response.(*model.UserModel)
		if response.Id != nil {
			err := u.repo.Put(response)
			if err != nil {
				http.Error(w, "DB error", http.StatusInternalServerError)
			} else {
				w.WriteHeader(200)
				json.NewEncoder(w).Encode(response)
			}
		} else {
			http.Error(w, "JSON parsing error", http.StatusBadRequest)
		}
	} else {
		http.Error(w, "JSON parsing error", http.StatusBadRequest)
	}
}

func (u user) Get(w http.ResponseWriter, req *http.Request) {
	response, err := u.repo.Get()
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
	} else {
		w.WriteHeader(200)
		if response == nil {
			response = []model.UserModel{}
		}
		json.NewEncoder(w).Encode(response)
	}

}
