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
	router.HandleFunc("/api/items", putItem).Methods("PUT")
	router.HandleFunc("/api/items", getItem).Methods("GET")

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
