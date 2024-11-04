package main

import (
	"HomeAssist/database"
	"HomeAssist/models"

	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

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

	database.SqlPing()
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/items", putItem).Methods("PUT")
	router.HandleFunc("/api/items", getItem).Methods("GET")

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
