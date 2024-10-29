package main

import (
	"log"
	"net/http"
	paths "stratplusapi/api"
	handler "stratplusapi/api/handler"

	"github.com/gorilla/mux"
)

func main() {

	mux := mux.NewRouter()

	mux.HandleFunc(paths.Login, handler.LoginHandler).Methods("POST")
	mux.HandleFunc(paths.CreateUser, handler.CreateUserHandler).Methods("POST")

	log.Println("Server started at 'localhost:8080'")
	http.ListenAndServe(":8080", mux)
}
