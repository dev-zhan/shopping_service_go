package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"shopping_service_go/controllers"
)

func main() {
	router := mux.NewRouter()

	//Create
	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")

	//Read
	router.HandleFunc("/users/{userId}", controllers.GetUser).Methods("GET")

	//Read-All
	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")

	//Update
	router.HandleFunc("/users/{userId}", controllers.UpdateUser).Methods("PUT")

	//Delete
	router.HandleFunc("/users/{userId}", controllers.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
