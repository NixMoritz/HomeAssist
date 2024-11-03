package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Message struct {
	Content string `json:"content"`
}

func getMessage(w http.ResponseWriter, r *http.Request) {
	response := Message{Content: "Hello from Go Backend!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func putMessage(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is PUT
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON body into a Message struct
	var msg Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Display the received message content in the console
	fmt.Println("Received message:", msg.Content)

	// Send a response back to the client
	w.Header().Set("Content-Type", "application/json")
	response := Message{Content: "Message received successfully"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/message", getMessage).Methods("GET")
	router.HandleFunc("/api/message", putMessage).Methods("PUT")

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

curl -X PUT http://localhost:8080/api/message -H "Content-Type: application/json" -d "{\"content\": \"Hello from the client!\"}"


