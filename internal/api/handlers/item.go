package handlers

import (
	"HomeAssist/internal/models"
	"HomeAssist/internal/storage/database"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

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

	router.HandleFunc("/api/items/create-batch", func(w http.ResponseWriter, r *http.Request) {
		putItems(w, r, db)
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
		log.Printf("Error decoding item: %v", err)
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Calculate price per unit
	if item.Units > 0 {
		item.Price_Per_Unit = item.Item_Price / item.Units
	} else {
		item.Price_Per_Unit = item.Item_Price
	}

	newItem, err := database.AddNewItem(item, db)
	if err != nil {
		log.Printf("Error adding item: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Error adding item")
		return
	}

	respondWithJSON(w, http.StatusCreated, newItem)
}

func putItems(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	var items models.Items
	if err := json.NewDecoder(r.Body).Decode(&items); err != nil {
		log.Printf("Error decoding items: %v", err)
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	var addedItems []models.Item
	for i, item := range items.Items {
		newItem, err := database.AddNewItem(item, db)
		if err != nil {
			log.Printf("Error adding item %d: %v", i+1, err)
			respondWithError(w, http.StatusInternalServerError, "Error adding item")
			return
		}

		addedItems = append(addedItems, newItem)
	}

	respondWithJSON(w, http.StatusCreated, map[string]interface{}{
		"message": "Items created successfully",
		"items":   addedItems,
	})
}

func updateItem(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		log.Printf("Error decoding item: %v", err)
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Set the updated timestamp
	item.Updated_At = time.Now()

	err := database.UpdateItem(item, db)
	if err != nil {
		log.Printf("Error updating item: %v", err)
		if strings.Contains(err.Error(), "not found") {
			respondWithError(w, http.StatusNotFound, err.Error())
		} else {
			respondWithError(w, http.StatusInternalServerError, "Error updating item")
		}
		return
	}

	respondWithJSON(w, http.StatusOK, item)
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
		log.Printf("Error getting items: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	respondWithJSON(w, http.StatusOK, items)
}
