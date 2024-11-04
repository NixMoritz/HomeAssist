package database

import (
	"HomeAssist/models"
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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

	columns := []string{}
	values := []interface{}{}
	valuePlaceholders := []string{}

	if item.ItemName != "" {
		columns = append(columns, "ItemName")
		values = append(values, item.ItemName)
		valuePlaceholders = append(valuePlaceholders, "$"+strconv.Itoa(len(values)))
	}
	if item.UnitPrice != 0 {
		columns = append(columns, "UnitPrice")
		values = append(values, item.UnitPrice)
		valuePlaceholders = append(valuePlaceholders, "$"+strconv.Itoa(len(values)))
	}
	if item.Units != 0 {
		columns = append(columns, "Units")
		values = append(values, item.Units)
		valuePlaceholders = append(valuePlaceholders, "$"+strconv.Itoa(len(values)))
	}
	if item.VendorID != 0 {
		columns = append(columns, "VendorID")
		values = append(values, item.VendorID)
		valuePlaceholders = append(valuePlaceholders, "$"+strconv.Itoa(len(values)))
	}
	if item.Weight != 0 {
		columns = append(columns, "Weight")
		values = append(values, item.Weight)
		valuePlaceholders = append(valuePlaceholders, "$"+strconv.Itoa(len(values)))
	}

	// Construct the final SQL statement
	insertSQL := "INSERT INTO items (" + strings.Join(columns, ", ") + ") VALUES (" + strings.Join(valuePlaceholders, ", ") + ")"

	// Execute the SQL statement with the collected values
	_, err := db.Exec(insertSQL, values...)
	if err != nil {
		log.Fatalf("Error inserting new item: %v", err)
	}
}

func Migration(db *sql.DB, migrationsDir string) {
	// Read all files in the migrations directory
	err := filepath.Walk(migrationsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the file is a regular file and ends with .sql
		if !info.IsDir() && filepath.Ext(path) == ".sql" {
			// Read the SQL file
			sqlFile, err := os.ReadFile(path)
			if err != nil {
				log.Printf("Error reading SQL file %s: %v", path, err)
				return err
			}

			// Split SQL commands if there are multiple statements
			commands := strings.Split(string(sqlFile), ";")
			for _, command := range commands {
				// Trim spaces and check if the command is not empty
				if trimmedCommand := strings.TrimSpace(command); trimmedCommand != "" {
					// Execute the SQL statement
					_, err := db.Exec(trimmedCommand)
					if err != nil {
						log.Printf("Error executing SQL command from file %s: %v", path, err)
					} else {
						// Log success
						log.Printf("Successfully executed command from %s: %s", path, trimmedCommand)
					}
				}
			}
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Error walking through migrations directory: %v", err)
	}
}
