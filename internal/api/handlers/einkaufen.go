package handlers

import (
	"HomeAssist/internal/models"
	"HomeAssist/internal/storage/database"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterEinkaufenHandlers(router *mux.Router, db *sql.DB) {
	router.HandleFunc("/einkaufen", func(w http.ResponseWriter, r *http.Request) {
		einkaufen(w, r, db)
	}).Methods(http.MethodPut)
}

func einkaufen(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var receiptEntry models.ReceiptEntry
	if err := json.NewDecoder(r.Body).Decode(&receiptEntry); err != nil {
		respondWithError(w, http.StatusBadRequest, "Bad request: invalid JSON format")
		return
	}

	database.AddEinkauf(receiptEntry, db)

	respondWithJSON(w, http.StatusCreated, map[string]string{"message": "ReceiptItem received successfully"})
}
