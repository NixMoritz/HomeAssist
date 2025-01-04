package handlers

import (
	"HomeAssist/internal/models"
	"HomeAssist/internal/storage/database"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterEinkaufenHandlers(router *mux.Router, db *sql.DB) {
	router.HandleFunc("/einkaufen", func(w http.ResponseWriter, r *http.Request) {
		einkaufen(w, r, db)
	}).Methods("PUT")
}

func einkaufen(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var receiptEntry models.ReceiptEntry
	if err := json.NewDecoder(r.Body).Decode(&receiptEntry); err != nil {
		http.Error(w, "Bad request: invalid JSON format", http.StatusBadRequest)
		return
	}

	fmt.Printf("Received receiptItem: %+v\n", receiptEntry)
	database.AddEinkauf(receiptEntry, db)

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "ReceiptItem received successfully"}
	json.NewEncoder(w).Encode(response)
}
