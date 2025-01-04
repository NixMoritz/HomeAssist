package handlers

import (
	"HomeAssist/internal/models"
	"HomeAssist/internal/storage/database"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func RegisterReceiptItemHandlers(router *mux.Router, db *sql.DB) {
	router.HandleFunc("/api/receiptItem", func(w http.ResponseWriter, r *http.Request) {
		updateItem(w, r, db)
	}).Methods("PUT")
	router.HandleFunc("/api/receiptItem", func(w http.ResponseWriter, r *http.Request) {
		updateItem(w, r, db)
	}).Methods("GET")
	router.HandleFunc("/api/receiptItem/all", func(w http.ResponseWriter, r *http.Request) {
		updateItem(w, r, db)
	}).Methods("GET")
}

func putReceiptItem(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var receiptItem models.ReceiptItem
	if err := json.NewDecoder(r.Body).Decode(&receiptItem); err != nil {
		http.Error(w, "Bad request: invalid JSON format", http.StatusBadRequest)
		return
	}

	fmt.Printf("Received receiptItem: %+v\n", receiptItem)
	database.AddNewReceiptItem(receiptItem, db)

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "ReceiptItem received successfully"}
	json.NewEncoder(w).Encode(response)
}

func getReceiptItem(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	receiptItemID, err := strconv.Atoi(r.URL.Query().Get("receiptItem_id"))
	if err != nil {
		http.Error(w, "Invalid receiptItem ID", http.StatusBadRequest)
		return
	}

	receiptItem, err := database.GetReceiptItem(receiptItemID, db)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(receiptItem)
}

func getAllReceiptItems(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	receiptItems, err := database.GetAllReceiptItems(db)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(receiptItems)
}
