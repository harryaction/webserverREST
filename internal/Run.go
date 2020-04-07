package internal

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"webserverREST/datasource"
	"webserverREST/internal/web/controllers"
)

func Run() {
	datasource.MustNewDB()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users", controllers.GetUser).Methods("GET")
	router.HandleFunc("/user/{id}", controllers.DropUser).Methods("DELETE")
	router.HandleFunc("/user", controllers.EditUser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", router))
}
