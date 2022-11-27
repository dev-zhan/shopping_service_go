package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"shopping_service_go/controllers"
	"shopping_service_go/database"
)

func main() {
	database.Connect("host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai")
	database.Migrate()
	router := mux.NewRouter()

	//Routing for Users
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

	//Routing for Items
	//Create
	router.HandleFunc("/items", controllers.CreateItem).Methods("POST")

	//Read
	router.HandleFunc("/items/{itemId}", controllers.GetItem).Methods("GET")

	//Read-All
	router.HandleFunc("/items", controllers.GetItems).Methods("GET")

	//Update
	router.HandleFunc("/items/{itemId}", controllers.UpdateItem).Methods("PUT")

	//Delete
	router.HandleFunc("/items/{itemId}", controllers.DeleteItem).Methods("DELETE")

	log.Println("...Starting on the server :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
