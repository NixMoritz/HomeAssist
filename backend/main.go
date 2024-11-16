package main

import (
	"HomeAssist/database"
	"HomeAssist/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Global variable to store the DB connection
var db *sql.DB

func putItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var items []models.Item
	if err := json.NewDecoder(r.Body).Decode(&items); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	for _, item := range items {
		fmt.Printf("Received item: %+v\n", item)
		// Pass the initialized db instance to AddNewItem
		database.AddNewItem(item, db)
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "Items received successfully"}
	json.NewEncoder(w).Encode(response)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

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

	// Run migrations using the initialized db instance
	migrationsDir := "../SQL Migration" // Adjust this path as needed
	database.Migration(db, migrationsDir)

	// Initialize router
	router := mux.NewRouter()
	router.HandleFunc("/hc", hc).Methods("GET")
	router.HandleFunc("/api/items", putItem).Methods("PUT")
	router.HandleFunc("/api/items", getItem).Methods("GET")

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
