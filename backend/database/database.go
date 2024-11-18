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

func AddNewStore(store models.Store, db *sql.DB) {

	// Execute the SQL statement with the collected values
	_, err := db.Exec(
		InsertStoreQuery,
		store.Unique_UID,
		store.Store_Branch,
		store.Store_Name,
		store.Store_Address,
		store.Store_Phone,
	)

	if err != nil {
		log.Fatalf("Error inserting new item: %v", err)
	}
}

func GetStore(storeID int, db *sql.DB) (*models.Store, error) {
	var store models.Store

	row := db.QueryRow(getStore, storeID)

	err := row.Scan(
		&store.Store_ID,
		&store.Unique_UID,
		&store.Store_Branch,
		&store.Store_Name,
		&store.Store_Address,
		&store.Store_Phone,
		&store.Updated_At,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("store with ID %d not found", storeID)
		}
		return nil, fmt.Errorf("error retrieving store: %w", err)
	}

	return &store, nil
}

func GetAllStores(db *sql.DB) ([]*models.Store, error) {
	var stores []*models.Store

	rows, err := db.Query(getAllStores)
	if err != nil {
		return nil, fmt.Errorf("error retrieving stores: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var store models.Store
		err := rows.Scan(
			&store.Store_ID,
			&store.Unique_UID,
			&store.Store_Branch,
			&store.Store_Name,
			&store.Store_Address,
			&store.Store_Phone,
			&store.Updated_At,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning store: %w", err)
		}
		stores = append(stores, &store)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating through stores: %w", err)
	}

	if len(stores) == 0 {
		return nil, fmt.Errorf("no stores found")
	}

	return stores, nil
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

func GetAllItems(db *sql.DB) ([]*models.Item, error) {
	var items []*models.Item

	rows, err := db.Query(getAllItems)
	if err != nil {
		return nil, fmt.Errorf("error retrieving items: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var item models.Item
		err := rows.Scan(
			&item.Item_ID,
			&item.Item_Name,
			&item.Unit_Price,
			&item.Units,
			&item.Store_Branch,
			&item.Weight,
			&item.Updated_At,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning item: %w", err)
		}
		items = append(items, &item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating through items: %w", err)
	}

	return items, nil
}
