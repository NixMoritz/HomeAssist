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

func RegisterReceiptHandlers(router *mux.Router, db *sql.DB) {
	router.HandleFunc("/api/receipt", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			putReceipt(w, r, db)
		case http.MethodGet:
			if r.URL.Query().Get("receipt_id") != "" {
				getReceipt(w, r, db)
			} else {
				getAllReceipts(w, r, db)
			}
		default:
			respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	})
}

func putReceipt(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	var receipt models.Receipt
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		respondWithError(w, http.StatusBadRequest, "Bad request: invalid JSON format")
		return
	}

	database.AddNewReceipt(receipt, db)

	respondWithJSON(w, http.StatusCreated, map[string]string{"message": "Receipt received successfully"})
}

func getReceipt(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	receiptID, err := strconv.Atoi(r.URL.Query().Get("receipt_id"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid receipt ID")
		return
	}

	receipt, err := database.GetReceipt(receiptID, db)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			respondWithError(w, http.StatusNotFound, err.Error())
		} else {
			respondWithError(w, http.StatusInternalServerError, "Internal server error")
		}
		return
	}

	respondWithJSON(w, http.StatusOK, receipt)
}

func getAllReceipts(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	receipts, err := database.GetAllReceipts(db)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			respondWithError(w, http.StatusNotFound, err.Error())
		} else {
			respondWithError(w, http.StatusInternalServerError, "Internal server error")
		}
		return
	}
	respondWithJSON(w, http.StatusOK, receipts)
}
