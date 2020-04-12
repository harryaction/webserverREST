package internal

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"webserverREST/datasource"
	"webserverREST/internal/repositories"
	"webserverREST/internal/web/controllers"
	"webserverREST/internal/web/handlers"
)

const (
	DB_HOST     = "178.62.8.121"
	DB_USER     = "alivedbuser"
	DB_PASSWORD = "7avsd84Egs_awNS4DXV"
	DB_NAME     = "thriatlonstore"
	WWW_PORT    = ":8080"
)

type Api struct {
	Router     *mux.Router
	Controller controllers.User
	DB         *sqlx.DB
}

func (a *Api) Setup() {
	if a.DB == nil {
		a.DB = datasource.MustNewDB()
	}
	repo := repositories.NewUser(a.DB)
	a.Controller = controllers.NewUser(repo)
	a.Router = mux.NewRouter().StrictSlash(true)
	a.Router.Use(handlers.Parse)
	a.Router.HandleFunc("/users", a.Controller.Create).Methods("POST")
	a.Router.HandleFunc("/users", a.Controller.Get).Methods("GET")
	a.Router.HandleFunc("/users/{id}", a.Controller.Delete).Methods("DELETE")
	a.Router.HandleFunc("/users", a.Controller.Edit).Methods("PUT")
}

func (a *Api) Run() {
	log.Fatalln(http.ListenAndServe(":8080", a.Router))
}

func Run() {
	a := Api{}
	a.Setup()
	a.Run()
}
