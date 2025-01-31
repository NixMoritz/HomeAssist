package database

import (
	"HomeAssist/internal/config"
	"HomeAssist/internal/models"
	queries "HomeAssist/internal/storage/HomeAssist"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var err error

const (
	MaxOpenConns    = 25
	MaxIdleConns    = 25
	ConnMaxLifetime = 5 * time.Minute
)

// InitDB initializes and returns a DB connection
func InitDB() (*sql.DB, error) {
	dbConfig := config.GetDatabaseValues()

	if dbConfig.Host == "" || dbConfig.Port == 0 || dbConfig.Username == "" || dbConfig.Password == "" || dbConfig.DBname == "" {
		return nil, fmt.Errorf("incomplete database configuration: %+v", dbConfig)
	}

	psqlInfo, err := buildConnectionString(*dbConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to build database connection string: %w", err)
	}

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	db.SetMaxOpenConns(MaxOpenConns)
	db.SetMaxIdleConns(MaxIdleConns)
	db.SetConnMaxLifetime(ConnMaxLifetime)

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Database connection successfully initialized with settings:")
	log.Printf("Host: %s, Port: %d, DB: %s\n", dbConfig.Host, dbConfig.Port, dbConfig.DBname)

	return db, nil
}

func buildConnectionString(cfg config.Database) (string, error) {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBname,
	), nil
}

func AddNewStore(store models.Store, db *sql.DB) {

	_, err := db.Exec(
		queries.InsertStoreQuery,
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

	row := db.QueryRow(queries.GetStore, storeID)

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

	rows, err := db.Query(queries.GetAllStores)
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

	_, err := db.Exec(
		queries.InsertItemQuery,
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

	row := db.QueryRow(queries.GetItem, itemID)

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

func DeleteItem(itemID int, db *sql.DB) error {
	result, err := db.Exec(queries.DeleteItem, itemID)
	if err != nil {
		return fmt.Errorf("error deleting item: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("item with ID %d not found", itemID)
	}

	return nil
}

func UpdateItem(item models.Item, db *sql.DB) error {
	fmt.Printf("Updating item with ID: %d\n", item.Item_ID)
	fmt.Printf("Values: Name=%s, Price=%f, Units=%f, Branch=%s, Weight=%f\n", item.Item_Name, item.Unit_Price, item.Units, item.Store_Branch, item.Weight)

	result, err := db.Exec(queries.UpdateItem, item.Item_Name, item.Unit_Price, item.Units, item.Store_Branch, item.Weight, item.Item_ID)
	if err != nil {
		return fmt.Errorf("error updating item: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("item with ID %d not found", item.Item_ID)
	}

	return nil
}

func GetAllItems(db *sql.DB) ([]*models.Item, error) {
	var items []*models.Item

	rows, err := db.Query(queries.GetAllItems)
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

func AddNewReceipt(receipt models.Receipt, db *sql.DB) {

	_, err := db.Exec(
		queries.InsertReceiptQuery,
		receipt.Store_ID,
		receipt.Date_Issued,
		receipt.Total_Amount,
		receipt.Payment_Method,
		receipt.Total_Discount_Amount,
		receipt.Net_Amount,
		receipt.Notes,
	)

	if err != nil {
		log.Fatalf("Error inserting new receipt: %v", err)
	}
}

func GetReceipt(receiptID int, db *sql.DB) (*models.Receipt, error) {
	var receipt models.Receipt

	row := db.QueryRow(queries.GetReceipt, receiptID)

	err := row.Scan(
		&receipt.Receipt_ID,
		&receipt.Store_ID,
		&receipt.Date_Issued,
		&receipt.Total_Amount,
		&receipt.Payment_Method,
		&receipt.Total_Discount_Amount,
		&receipt.Net_Amount,
		&receipt.Notes,
		&receipt.Updated_At,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("receipt with ID %d not found", receiptID)
		}
		return nil, fmt.Errorf("error retrieving receipt: %w", err)
	}

	return &receipt, nil
}

func GetAllReceipts(db *sql.DB) ([]*models.Receipt, error) {
	var receipts []*models.Receipt

	rows, err := db.Query(queries.GetAllReceipts)
	if err != nil {
		return nil, fmt.Errorf("error retrieving receipts: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var receipt models.Receipt
		err := rows.Scan(
			&receipt.Receipt_ID,
			&receipt.Store_ID,
			&receipt.Date_Issued,
			&receipt.Total_Amount,
			&receipt.Payment_Method,
			&receipt.Total_Discount_Amount,
			&receipt.Net_Amount,
			&receipt.Notes,
			&receipt.Updated_At,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning receipt: %w", err)
		}
		receipts = append(receipts, &receipt)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating through receipts: %w", err)
	}

	return receipts, nil
}

func AddNewReceiptItem(receiptItem models.ReceiptItem, db *sql.DB) {

	_, err := db.Exec(
		queries.InsertReceiptItemQuery,
		receiptItem.Receipt_ID,
		receiptItem.Item_ID,
		receiptItem.Quantity,
		receiptItem.Discount_Amount,
		receiptItem.Total_Price,
	)

	if err != nil {
		log.Fatalf("Error inserting new receiptItem: %v", err)
	}
}

func GetReceiptItem(receiptItemID int, db *sql.DB) (*models.ReceiptItem, error) {
	var receiptItem models.ReceiptItem

	row := db.QueryRow(queries.GetReceiptItem, receiptItemID)

	err := row.Scan(
		&receiptItem.Receipt_Item_ID,
		&receiptItem.Receipt_ID,
		&receiptItem.Item_ID,
		&receiptItem.Quantity,
		&receiptItem.Discount_Amount,
		&receiptItem.Total_Price,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("receiptItem with ID %d not found", receiptItemID)
		}
		return nil, fmt.Errorf("error retrieving receiptItem: %w", err)
	}

	return &receiptItem, nil
}

func GetAllReceiptItems(db *sql.DB) ([]*models.ReceiptItem, error) {
	var receiptItems []*models.ReceiptItem

	rows, err := db.Query(queries.GetAllReceiptItems)
	if err != nil {
		return nil, fmt.Errorf("error retrieving receiptItems: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var receiptItem models.ReceiptItem
		err := rows.Scan(
			&receiptItem.Receipt_Item_ID,
			&receiptItem.Receipt_ID,
			&receiptItem.Item_ID,
			&receiptItem.Quantity,
			&receiptItem.Discount_Amount,
			&receiptItem.Total_Price,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning receiptItem: %w", err)
		}
		receiptItems = append(receiptItems, &receiptItem)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating through receiptItems: %w", err)
	}

	return receiptItems, nil
}

func AddEinkauf(receiptEntry models.ReceiptEntry, db *sql.DB) {

	fmt.Println(receiptEntry.Receipt)
	for i := 0; i < len(receiptEntry.ReceiptItems); i++ {
		fmt.Println(receiptEntry.ReceiptItems[i])
	}

	AddNewReceipt(receiptEntry.Receipt, db)
	for i := 0; i < len(receiptEntry.ReceiptItems); i++ {
		AddNewReceiptItem(receiptEntry.ReceiptItems[i], db)
	}
}
