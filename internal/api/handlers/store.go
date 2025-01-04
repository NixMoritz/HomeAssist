// store.go
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

func RegisterStoreHandlers(router *mux.Router, db *sql.DB) {
	router.HandleFunc("/api/stores", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			putStore(w, r, db)
		} else {
			getStore(w, r, db)
		}
	}).Methods("PUT", "GET")

	router.HandleFunc("/api/stores/all", func(w http.ResponseWriter, r *http.Request) {
		getAllStores(w, r, db)
	}).Methods("GET")
}

func putStore(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var store models.Store
	if err := json.NewDecoder(r.Body).Decode(&store); err != nil {
		http.Error(w, "Bad request: invalid JSON format", http.StatusBadRequest)
		return
	}
	database.AddNewStore(store, db)
	response := map[string]string{"message": "store received successfully"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getStore(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	storeID, err := strconv.Atoi(r.URL.Query().Get("store_id"))
	if err != nil {
		http.Error(w, "Invalid store ID", http.StatusBadRequest)
		return
	}

	store, err := database.GetStore(storeID, db)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(store)
}

func getAllStores(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	stores, err := database.GetAllStores(db)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stores)
}
