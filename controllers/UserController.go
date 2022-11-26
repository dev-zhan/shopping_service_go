package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"shopping_service_go/models"
	"strconv"
)

var (
	users      []models.User
	prevUserID = 0
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	prevUserID++
	user.UserID = strconv.Itoa(prevUserID)
	users = append(users, user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
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

func UpdateUser(w http.ResponseWriter, r *http.Request) {
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

func DeleteUser(w http.ResponseWriter, r *http.Request) {
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
