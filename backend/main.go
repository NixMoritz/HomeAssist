package main

import (
	"HomeAssist/database"
	"HomeAssist/models"
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

func hc(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	database.Healthcheck(db)
}

func main() {
	var err error
	db, err = database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	migrationsDir := "../SQL Migration"
	database.Migration(db, migrationsDir)

	router := mux.NewRouter()
	router.HandleFunc("/hc", hc).Methods("GET")
	router.HandleFunc("/api/stores", putStore).Methods("PUT")
	router.HandleFunc("/api/stores", getStore).Methods("GET")
	router.HandleFunc("/api/stores/all", getAllStores).Methods("GET")

	router.HandleFunc("/api/items", putItem).Methods("PUT")
	router.HandleFunc("/api/items", getItem).Methods("GET")
	router.HandleFunc("/api/items/all", getAllItems).Methods("GET")

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
