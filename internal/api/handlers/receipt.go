// receipt.go
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

func RegisterReceiptHandlers(router *mux.Router, db *sql.DB) {
	router.HandleFunc("/api/receipt", func(w http.ResponseWriter, r *http.Request) {
		putReceipt(w, r, db)
	}).Methods("PUT")
	router.HandleFunc("/api/receipt", func(w http.ResponseWriter, r *http.Request) {
		getReceipt(w, r, db)
	}).Methods("GET")
	router.HandleFunc("/api/receipt", func(w http.ResponseWriter, r *http.Request) {
		getAllReceipts(w, r, db)
	}).Methods("GET")
}
func putReceipt(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var receipt models.Receipt
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		http.Error(w, "Bad request: invalid JSON format", http.StatusBadRequest)
		return
	}

	fmt.Printf("Received receipt: %+v\n", receipt)
	database.AddNewReceipt(receipt, db)

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "Receipt received successfully"}
	json.NewEncoder(w).Encode(response)
}

func getReceipt(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	receiptID, err := strconv.Atoi(r.URL.Query().Get("receipt_id"))
	if err != nil {
		http.Error(w, "Invalid receipt ID", http.StatusBadRequest)
		return
	}

	receipt, err := database.GetReceipt(receiptID, db)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(receipt)
}

func getAllReceipts(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	receipts, err := database.GetAllReceipts(db)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(receipts)
}
