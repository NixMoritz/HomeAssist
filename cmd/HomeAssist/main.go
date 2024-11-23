package main

import (
	"HomeAssist/internal/api/HomeAssist"
	"HomeAssist/internal/models"
	"HomeAssist/internal/storage/database"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// Global variable to store the DB connection
var db *sql.DB

func putStore(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var store models.Store
	if err := json.NewDecoder(r.Body).Decode(&store); err != nil {
		http.Error(w, "Bad request: invalid JSON format", http.StatusBadRequest)
		return
	}

	fmt.Printf("Received store: %+v\n", store)
	database.AddNewStore(store, db)

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "store received successfully"}
	json.NewEncoder(w).Encode(response)
}

func getStore(w http.ResponseWriter, r *http.Request) {
	storeID, err := strconv.Atoi(r.URL.Query().Get("store_id"))
	if err != nil {
		http.Error(w, "Invalid store ID", http.StatusBadRequest)
		return
	}

	store, err := database.GetStore(storeID, db)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(store)
}

func getAllStores(w http.ResponseWriter, r *http.Request) {
	stores, err := database.GetAllStores(db)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stores)
}

func putItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Bad request: invalid JSON format", http.StatusBadRequest)
		return
	}

	fmt.Printf("Received item: %+v\n", item)
	database.AddNewItem(item, db)

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "Item received successfully"}
	json.NewEncoder(w).Encode(response)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	itemID, err := strconv.Atoi(r.URL.Query().Get("item_id"))
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	item, err := database.GetItem(itemID, db)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func getAllItems(w http.ResponseWriter, r *http.Request) {
	items, err := database.GetAllItems(db)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func putReceipt(w http.ResponseWriter, r *http.Request) {
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

func getReceipt(w http.ResponseWriter, r *http.Request) {
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

func getAllReceipts(w http.ResponseWriter, r *http.Request) {
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

func putReceiptItem(w http.ResponseWriter, r *http.Request) {
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

func getReceiptItem(w http.ResponseWriter, r *http.Request) {
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

func getAllReceiptItems(w http.ResponseWriter, r *http.Request) {
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

func hc(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	HomeAssist.GetHC(w, db)
}

func einkaufen(w http.ResponseWriter, r *http.Request) {
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

func main() {
	var err error
	db, err = database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	database.Migration(db)

	router := mux.NewRouter()
	router.HandleFunc("/hc", hc).Methods("GET")

	router.HandleFunc("/api/stores", putStore).Methods("PUT")
	router.HandleFunc("/api/stores", getStore).Methods("GET")
	router.HandleFunc("/api/stores/all", getAllStores).Methods("GET")

	router.HandleFunc("/api/items", putItem).Methods("PUT")
	router.HandleFunc("/api/items", getItem).Methods("GET")
	router.HandleFunc("/api/items/all", getAllItems).Methods("GET")

	router.HandleFunc("/api/receipt", putReceipt).Methods("PUT")
	router.HandleFunc("/api/receipt", getReceipt).Methods("GET")
	router.HandleFunc("/api/receipt/all", getAllReceipts).Methods("GET")

	router.HandleFunc("/api/receiptItem", putReceiptItem).Methods("PUT")
	router.HandleFunc("/api/receiptItem", getReceiptItem).Methods("GET")
	router.HandleFunc("/api/receiptItem/all", getAllReceiptItems).Methods("GET")

	router.HandleFunc("/einkaufen", einkaufen).Methods("PUT")

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
