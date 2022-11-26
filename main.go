package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"shopping_service_go/models"
	"strconv"
)

var (
	users      []models.User
	prevUserID = 0
)

func main() {
	router := mux.NewRouter()

	//Create
	router.HandleFunc("/users", createUser).Methods("POST")

	//Read
	router.HandleFunc("/users/{userId}", getUser).Methods("GET")

	//Read-All
	router.HandleFunc("/users", getUsers).Methods("GET")

	//Update
	router.HandleFunc("/users/{userId}", updateUser).Methods("PUT")

	//Delete
	router.HandleFunc("/users/{userId}", deleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	prevUserID++
	user.UserID = strconv.Itoa(prevUserID)
	users = append(users, user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputUserID := params["userId"]
	for _, user := range users {
		if user.UserID == inputUserID {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputUserID := params["userId"]
	for i, user := range users {
		if user.UserID == inputUserID {
			users = append(users[:i], users[i+1:]...)
			var updateUser models.User
			json.NewDecoder(r.Body).Decode(&updateUser)
			users = append(users, updateUser)
			json.NewEncoder(w).Encode(updateUser)
			return
		}
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputUserID := params["userId"]
	for i, user := range users {
		if user.UserID == inputUserID {
			users = append(users[:i], users[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
}
