package database

import (
	"HomeAssist/models"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "admin"
	dbname   = "HomeAssist"
)

var err error

// InitDB initializes and returns a DB connection
func InitDB() (*sql.DB, error) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open a database connection
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	// Set database connection pool settings
	db.SetMaxOpenConns(25)                 // Maximum number of open connections to the database
	db.SetMaxIdleConns(25)                 // Maximum number of connections in the idle connection pool
	db.SetConnMaxLifetime(5 * time.Minute) // Maximum amount of time a connection may be reused

	// Test the database connection
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	log.Println("Database connection successfully initialized")
	return db, nil
}
func Healthcheck(db *sql.DB) {
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	fmt.Println("Successfully connected!")
}

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

func AddNewItem(item models.Item, db *sql.DB) {

	// Execute the SQL statement with the collected values
	_, err := db.Exec(
		InsertItemQuery,
		item.Item_Name,
		item.Unit_Price,
		item.Units,
		item.Store_Branch,
		item.Weight,
	)

	if err != nil {
		log.Fatalf("Error inserting new item: %v", err)
	}
}

func GetItem(itemID int, db *sql.DB) (*models.Item, error) {
	var item models.Item

	row := db.QueryRow(getItem, itemID)

	err := row.Scan(
		&item.Item_ID,
		&item.Item_Name,
		&item.Unit_Price,
		&item.Units,
		&item.Store_Branch,
		&item.Weight,
		&item.Updated_At,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("item with ID %d not found", itemID)
		}
		return nil, fmt.Errorf("error retrieving item: %w", err)
	}

	return &item, nil
}
