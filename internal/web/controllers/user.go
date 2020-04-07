package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"webserverREST/internal/model"
	"webserverREST/internal/repositories"
	"webserverREST/internal/web/binders"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := binders.Parse(r, true, false)
	if err != nil {
		http.Error(w, "can't read body", http.StatusBadRequest)
	} else {
		response := model.UserModel{
			Id:        body.Id,
			Name:      body.Name,
			Lastname:  body.Lastname,
			Age:       body.Age,
			Birthdate: body.Birthdate,
		}
		u := repositories.NewUser()
		err := u.Put(response)
		if err != nil {
			w.Header().Set("Server", "REST SERVER")
			http.Error(w, "db error", http.StatusInternalServerError)
		} else {
			w.Header().Set("Server", "REST SERVER")
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(response)
		}
	}
}

func DropUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := vars["id"]
	u := repositories.NewUser()
	err := u.Delete(id)
	if err != nil {
		w.Header().Set("Server", "REST SERVER")
		http.Error(w, "db error", http.StatusInternalServerError)
	} else {
		w.Header().Set("Server", "REST SERVER")
		w.WriteHeader(200)
	}
}

func EditUser(w http.ResponseWriter, r *http.Request) {
	body, err := binders.Parse(r, false, true)
	if err != nil {
		http.Error(w, "can't read body", http.StatusBadRequest)
	} else {
		response := model.UserModel{
			Id:        body.Id,
			Name:      body.Name,
			Lastname:  body.Lastname,
			Age:       body.Age,
			Birthdate: body.Birthdate,
		}
		u := repositories.NewUser()
		err := u.Put(response)
		if err != nil {
			w.Header().Set("Server", "REST SERVER")
			http.Error(w, "db error", http.StatusInternalServerError)
		} else {
			w.Header().Set("Server", "REST SERVER")
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(response)
		}
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	u := repositories.NewUser()
	response, err := u.Get()
	if err != nil {
		w.Header().Set("Server", "REST SERVER")
		http.Error(w, "db error", http.StatusInternalServerError)
	} else {
		w.Header().Set("Server", "REST SERVER")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(response)
	}
}
