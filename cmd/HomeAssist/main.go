package main

import (
	"HomeAssist/internal/api/handlers"
	"HomeAssist/internal/api/middleware"
	"HomeAssist/internal/config"
	"HomeAssist/internal/storage/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	var err error

	config.LoadConfig()

	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	database.Migration(db)

	router := mux.NewRouter()

	// Register handlers
	handlers.RegisterStoreHandlers(router, db)
	handlers.RegisterItemHandlers(router, db)
	handlers.RegisterReceiptHandlers(router, db)
	handlers.RegisterReceiptItemHandlers(router, db)
	handlers.RegisterHealthCheckHandlers(router, db)
	handlers.RegisterEinkaufenHandlers(router, db)

	// Apply middleware
	cors := middleware.CorsMiddleware(router)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", cors))
}
