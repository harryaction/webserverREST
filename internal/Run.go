package internal

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"webserverREST/internal/datasource"
	"webserverREST/internal/repositories"
	"webserverREST/internal/web/controllers"
	"webserverREST/internal/web/handlers"
)

type Api struct {
	Router     *mux.Router
	Controller controllers.User
	DB         *sqlx.DB
}

func (a *Api) Setup() {
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
	db := datasource.MustNewDB()
	a := Api{DB: db}
	a.Setup()
	a.Run()
}
