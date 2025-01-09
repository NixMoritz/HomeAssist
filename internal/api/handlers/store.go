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
		switch r.Method {
		case http.MethodPut:
			putStore(w, r, db)
		case http.MethodGet:
			getStore(w, r, db)
		default:
			respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	})

	router.HandleFunc("/api/stores/all", func(w http.ResponseWriter, r *http.Request) {
		getAllStores(w, r, db)
	}).Methods(http.MethodGet)
}

func putStore(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var store models.Store
	if err := json.NewDecoder(r.Body).Decode(&store); err != nil {
		respondWithError(w, http.StatusBadRequest, "Bad request: invalid JSON format")
		return
	}
	database.AddNewStore(store, db)
	respondWithJSON(w, http.StatusCreated, map[string]string{"message": "Store created successfully"})
}

func getStore(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	storeID, err := strconv.Atoi(r.URL.Query().Get("store_id"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid store ID")
		return
	}

	store, err := database.GetStore(storeID, db)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			respondWithError(w, http.StatusNotFound, err.Error())
		} else {
			respondWithError(w, http.StatusInternalServerError, "Internal server error")
		}
		return
	}
	respondWithJSON(w, http.StatusOK, store)
}

func getAllStores(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	stores, err := database.GetAllStores(db)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			respondWithError(w, http.StatusNotFound, err.Error())
		} else {
			respondWithError(w, http.StatusInternalServerError, "Internal server error")
		}
		return
	}
	respondWithJSON(w, http.StatusOK, stores)
}
