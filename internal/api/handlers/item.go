package handlers

import (
	"HomeAssist/internal/models"
	"HomeAssist/internal/storage/database"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func RegisterItemHandlers(router *mux.Router, db *sql.DB) {
	router.HandleFunc("/api/items/delete", func(w http.ResponseWriter, r *http.Request) {
		deleteItem(w, r, db)
	}).Methods("DELETE")

	router.HandleFunc("/api/items/update", func(w http.ResponseWriter, r *http.Request) {
		updateItem(w, r, db)
	}).Methods("PUT")

	router.HandleFunc("/api/items/create", func(w http.ResponseWriter, r *http.Request) {
		putItem(w, r, db)
	}).Methods("PUT")

	router.HandleFunc("/api/items", func(w http.ResponseWriter, r *http.Request) {
		getItem(w, r, db)
	}).Methods("GET")

	router.HandleFunc("/api/items/all", func(w http.ResponseWriter, r *http.Request) {
		getAllItems(w, r, db)
	}).Methods("GET")
}

func putItem(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Bad request: invalid JSON format", http.StatusBadRequest)
		return
	}
	database.AddNewItem(item, db)
	response := map[string]string{"message": "Item received successfully"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func updateItem(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Bad request: invalid JSON format", http.StatusBadRequest)
		return
	}

	if item.Item_ID == 0 {
		http.Error(w, "Bad request: item ID is missing", http.StatusBadRequest)
		return
	}

	err := database.UpdateItem(item, db)
	if err != nil {
		http.Error(w, "Internal server error: could not update item", http.StatusInternalServerError)
		return
	}
	response := map[string]string{"message": "Item updated successfully"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func deleteItem(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	itemID, err := strconv.Atoi(r.URL.Query().Get("item_id"))
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	err = database.DeleteItem(itemID, db)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func getItem(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	itemID, err := strconv.Atoi(r.URL.Query().Get("item_id"))
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	item, err := database.GetItem(itemID, db)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func getAllItems(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	items, err := database.GetAllItems(db)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}
