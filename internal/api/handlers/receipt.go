package handlers

import (
	"HomeAssist/internal/models"
	"HomeAssist/internal/storage/database"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func RegisterReceiptHandlers(router *mux.Router, db *sql.DB) {
	router.HandleFunc("/api/receipts", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getAllReceipts(w, r, db)
		default:
			respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	})

	router.HandleFunc("/api/receipts/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
			return
		}
		createReceipt(w, r, db)
	})

	router.HandleFunc("/api/receipts/update", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
			return
		}
		updateReceipt(w, r, db)
	})

	router.HandleFunc("/api/receipts/delete", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
			return
		}
		deleteReceipt(w, r, db)
	})
}

func createReceipt(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var receiptInput struct {
		Store_ID              int     `json:"store_id"`
		Date_Issued           string  `json:"date_issued"`
		Total_Amount          float64 `json:"total_amount"`
		Payment_Method        string  `json:"payment_method"`
		Total_Discount_Amount float64 `json:"total_discount_amount"`
		Net_Amount            float64 `json:"net_amount"`
		Tax_Amount            float64 `json:"tax_amount"`
		Receipt_Type          string  `json:"receipt_type"`
		Loyalty_Card_Number   string  `json:"loyalty_card_number"`
		Cashier_Name          string  `json:"cashier_name"`
		Receipt_Number        string  `json:"receipt_number"`
		Notes                 string  `json:"notes"`
		Image_URL             string  `json:"image_url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&receiptInput); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Parse the ISO format date-time string
	parsedTime, err := time.Parse("2006-01-02T15:04", receiptInput.Date_Issued)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid date format: %v", err))
		return
	}

	receipt := models.Receipt{
		Store_ID:              receiptInput.Store_ID,
		Date_Issued:           parsedTime,
		Total_Amount:          receiptInput.Total_Amount,
		Payment_Method:        receiptInput.Payment_Method,
		Total_Discount_Amount: receiptInput.Total_Discount_Amount,
		Net_Amount:            receiptInput.Net_Amount,
		Tax_Amount:            receiptInput.Tax_Amount,
		Receipt_Type:          receiptInput.Receipt_Type,
		Loyalty_Card_Number:   receiptInput.Loyalty_Card_Number,
		Cashier_Name:          receiptInput.Cashier_Name,
		Receipt_Number:        receiptInput.Receipt_Number,
		Notes:                 receiptInput.Notes,
		Image_URL:             receiptInput.Image_URL,
	}

	if err := database.AddNewReceipt(receipt, db); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, receipt)
}

func getAllReceipts(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	receipts, err := database.GetAllReceipts(db)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, receipts)
}

func updateReceipt(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var receiptInput struct {
		Receipt_ID            int     `json:"receipt_id"`
		Store_ID              int     `json:"store_id"`
		Date_Issued           string  `json:"date_issued"`
		Time_Issued           string  `json:"time_issued"`
		Total_Amount          float64 `json:"total_amount"`
		Payment_Method        string  `json:"payment_method"`
		Total_Discount_Amount float64 `json:"total_discount_amount"`
		Net_Amount            float64 `json:"net_amount"`
		Tax_Amount            float64 `json:"tax_amount"`
		Receipt_Type          string  `json:"receipt_type"`
		Loyalty_Card_Number   string  `json:"loyalty_card_number"`
		Cashier_Name          string  `json:"cashier_name"`
		Receipt_Number        string  `json:"receipt_number"`
		Notes                 string  `json:"notes"`
		Image_URL             string  `json:"image_url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&receiptInput); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Log the received date string
	log.Printf("Received date string: %s", receiptInput.Date_Issued)

	// Parse the ISO format date-time string
	parsedTime, err := time.Parse("2006-01-02T15:04", receiptInput.Date_Issued)
	if err != nil {
		log.Printf("Date parsing error: %v", err)
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid date format. Expected YYYY-MM-DDThh:mm, got: %s", receiptInput.Date_Issued))
		return
	}

	// Create the final datetime
	dateTime := time.Date(
		parsedTime.Year(),
		parsedTime.Month(),
		parsedTime.Day(),
		parsedTime.Hour(),
		parsedTime.Minute(),
		0, // seconds
		0, // nanoseconds
		time.UTC,
	)

	receipt := models.Receipt{
		Receipt_ID:            receiptInput.Receipt_ID,
		Store_ID:              receiptInput.Store_ID,
		Date_Issued:           dateTime,
		Total_Amount:          receiptInput.Total_Amount,
		Payment_Method:        receiptInput.Payment_Method,
		Total_Discount_Amount: receiptInput.Total_Discount_Amount,
		Net_Amount:            receiptInput.Net_Amount,
		Tax_Amount:            receiptInput.Tax_Amount,
		Receipt_Type:          receiptInput.Receipt_Type,
		Loyalty_Card_Number:   receiptInput.Loyalty_Card_Number,
		Cashier_Name:          receiptInput.Cashier_Name,
		Receipt_Number:        receiptInput.Receipt_Number,
		Notes:                 receiptInput.Notes,
		Image_URL:             receiptInput.Image_URL,
	}

	if err := database.UpdateReceipt(receipt, db); err != nil {
		if strings.Contains(err.Error(), "not found") {
			respondWithError(w, http.StatusNotFound, err.Error())
			return
		}
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, receipt)
}

func deleteReceipt(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	receiptID, err := strconv.Atoi(r.URL.Query().Get("receipt_id"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid receipt ID")
		return
	}

	if err := database.DeleteReceipt(receiptID, db); err != nil {
		if strings.Contains(err.Error(), "not found") {
			respondWithError(w, http.StatusNotFound, err.Error())
			return
		}
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Receipt deleted successfully"})
}
