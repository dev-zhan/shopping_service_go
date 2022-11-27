package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"shopping_service_go/database"
	"shopping_service_go/models"
)

func CreateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item models.Item
	json.NewDecoder(r.Body).Decode(&item)
	database.Instance.Create(&item)

	json.NewEncoder(w).Encode(item)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputID := params["itemId"]
	if !checkIfUserExists(inputID) {
		json.NewEncoder(w).Encode("User Not Found!")
		return
	}
	var item models.Item
	database.Instance.First(&item, inputID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func checkIfItemExists(userId string) bool {
	var item models.Item
	database.Instance.First(&item, userId)
	if item.ID == 0 {
		return false
	}
	return true
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	var items []models.Item
	database.Instance.Find(&items)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	inputID := params["itemId"]
	if !checkIfUserExists(inputID) {
		json.NewEncoder(w).Encode("User Not Found!")
		return
	}
	var item models.Item
	database.Instance.First(&item, inputID)
	json.NewDecoder(r.Body).Decode(&item)
	database.Instance.Save(&item)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	inputID := params["itemId"]
	if !checkIfUserExists(inputID) {
		json.NewEncoder(w).Encode("User Not Found!")
		return
	}
	var item models.Item
	database.Instance.Delete(&item, inputID)
	json.NewEncoder(w).Encode("Deleted successfully!")
}
