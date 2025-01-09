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
	}).Methods(http.MethodDelete)

	router.HandleFunc("/api/items/update", func(w http.ResponseWriter, r *http.Request) {
		updateItem(w, r, db)
	}).Methods(http.MethodPut)

	router.HandleFunc("/api/items/create", func(w http.ResponseWriter, r *http.Request) {
		putItem(w, r, db)
	}).Methods(http.MethodPut)

	router.HandleFunc("/api/items", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("item_id") != "" {
			getItem(w, r, db)
		} else {
			getAllItems(w, r, db)
		}

	}).Methods(http.MethodGet)
}

func putItem(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		respondWithError(w, http.StatusBadRequest, "Bad request: invalid JSON format")
		return
	}
	database.AddNewItem(item, db)
	respondWithJSON(w, http.StatusCreated, map[string]string{"message": "Item created successfully"})
}

func updateItem(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		respondWithError(w, http.StatusBadRequest, "Bad request: invalid JSON format")
		return
	}

	if item.Item_ID == 0 {
		respondWithError(w, http.StatusBadRequest, "Bad request: item ID is missing")
		return
	}

	err := database.UpdateItem(item, db)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal server error: could not update item")
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Item updated successfully"})
}

func deleteItem(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	itemID, err := strconv.Atoi(r.URL.Query().Get("item_id"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid item ID")
		return
	}

	err = database.DeleteItem(itemID, db)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			respondWithError(w, http.StatusNotFound, err.Error())
		} else {
			respondWithError(w, http.StatusInternalServerError, "Internal server error")
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func getItem(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	itemID, err := strconv.Atoi(r.URL.Query().Get("item_id"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid item ID")
		return
	}

	item, err := database.GetItem(itemID, db)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			respondWithError(w, http.StatusNotFound, err.Error())
		} else {
			respondWithError(w, http.StatusInternalServerError, "Internal server error")
		}
		return
	}
	respondWithJSON(w, http.StatusOK, item)
}

func getAllItems(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	items, err := database.GetAllItems(db)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			respondWithError(w, http.StatusNotFound, err.Error())
		} else {
			respondWithError(w, http.StatusInternalServerError, "Internal server error")
		}
		return
	}

	respondWithJSON(w, http.StatusOK, items)
}
