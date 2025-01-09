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

func RegisterReceiptItemHandlers(router *mux.Router, db *sql.DB) {
	router.HandleFunc("/api/receiptItem", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			putReceiptItem(w, r, db)
		case http.MethodGet:
			if r.URL.Query().Get("receiptItem_id") != "" {
				getReceiptItem(w, r, db)
			} else {
				getAllReceiptItems(w, r, db)
			}
		default:
			respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	})

}

func putReceiptItem(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var receiptItem models.ReceiptItem
	if err := json.NewDecoder(r.Body).Decode(&receiptItem); err != nil {
		respondWithError(w, http.StatusBadRequest, "Bad request: invalid JSON format")
		return
	}

	database.AddNewReceiptItem(receiptItem, db)
	respondWithJSON(w, http.StatusCreated, map[string]string{"message": "ReceiptItem received successfully"})
}

func getReceiptItem(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	receiptItemID, err := strconv.Atoi(r.URL.Query().Get("receiptItem_id"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid receiptItem ID")
		return
	}

	receiptItem, err := database.GetReceiptItem(receiptItemID, db)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			respondWithError(w, http.StatusNotFound, err.Error())
		} else {
			respondWithError(w, http.StatusInternalServerError, "Internal server error")
		}
		return
	}

	respondWithJSON(w, http.StatusOK, receiptItem)
}

func getAllReceiptItems(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	receiptItems, err := database.GetAllReceiptItems(db)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			respondWithError(w, http.StatusNotFound, err.Error())
		} else {
			respondWithError(w, http.StatusInternalServerError, "Internal server error")
		}
		return
	}
	respondWithJSON(w, http.StatusOK, receiptItems)
}
