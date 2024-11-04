package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "admin"
	dbname   = "HomeAssist"
)

func SqlPing() {

	// Connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open connection to the database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	// Verify connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	fmt.Println("Successfully connected!")

	//Create a table
	createTableSQL := `CREATE TABLE Receipts (
    ReceiptID SERIAL PRIMARY KEY,               -- Unique identifier for each receipt
    ReceiptNumber VARCHAR(50) NOT NULL UNIQUE,  -- Unique receipt number for reference
    DateIssued DATE NOT NULL,                    -- Date when the receipt was issued
    CustomerName VARCHAR(255),                   -- Name of the customer
    CustomerEmail VARCHAR(255),                  -- Email address of the customer
    TotalAmount DECIMAL(10, 2) NOT NULL,         -- Total amount of the receipt
    PaymentMethod VARCHAR(50),                   -- Payment method used (e.g., Cash, Credit Card)
    TaxAmount DECIMAL(10, 2),                    -- Amount of tax applied
    DiscountAmount DECIMAL(10, 2),               -- Discount amount applied, if any
    NetAmount DECIMAL(10, 2) NOT NULL,           -- Final amount after applying tax and discount
    Notes TEXT,                                  -- Additional notes or comments
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp for when the receipt was created
    UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Timestamp for last update
);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}

	// Insert a new user
	insertSQL := `INSERT INTO users (name, age) VALUES ($1, $2)`
	_, err = db.Exec(insertSQL, "John Doe", 30)
	if err != nil {
		log.Fatalf("Error inserting new user: %v", err)
	}

	// Query the database
	var (
		id   int
		name string
		age  int
	)
	querySQL := `SELECT id, name, age FROM users WHERE name=$1`
	row := db.QueryRow(querySQL, "John Doe")
	err = row.Scan(&id, &name, &age)
	if err != nil {
		log.Fatalf("Error querying database: %v", err)
	}
	fmt.Printf("User found: %d, %s, %d\n", id, name, age)

}
