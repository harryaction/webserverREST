package internal

import (
	"github.com/gorilla/mux"
	"net/http"
	"webserverREST/datasource"
	"webserverREST/internal/repositories"
	"webserverREST/internal/web/controllers"
	"webserverREST/internal/web/handlers"
)

func Run() {
	datasource.MustNewDB()
	c := userController()

	router := mux.NewRouter().StrictSlash(true)
	router.Use(handlers.Parse)
	router.HandleFunc("/users", c.Create).Methods("POST")
	router.HandleFunc("/users", c.Get).Methods("GET")
	router.HandleFunc("/user/{id}", c.Delete).Methods("DELETE")
	router.HandleFunc("/user", c.Edit).Methods("PUT")
	http.ListenAndServe(":8080", router)
}

func userController() controllers.User {
	repo := repositories.NewUser()
	return controllers.NewUser(repo)
}
