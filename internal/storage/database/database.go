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

func AddNewItem(item models.Item, db *sql.DB) (models.Item, error) {
	_, err := db.Exec(
		queries.InsertItemQuery,
		item.Item_Name,
		item.Unit_Price,
		item.Units,
		item.Store_Branch,
		item.Weight,
		item.Category,
		item.Subcategory,
		item.Is_Organic,
		item.Brand_Name,
		item.Barcode,
	)

	if err != nil {
		return models.Item{}, fmt.Errorf("error inserting item: %w", err)
	}

	return item, nil
}

func GetItem(itemID int, db *sql.DB) (*models.Item, error) {
	var item models.Item
	err := db.QueryRow(queries.GetItem, itemID).Scan(
		&item.Item_ID,
		&item.Item_Name,
		&item.Unit_Price,
		&item.Units,
		&item.Store_Branch,
		&item.Weight,
		&item.Category,
		&item.Subcategory,
		&item.Is_Organic,
		&item.Brand_Name,
		&item.Barcode,
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
	result, err := db.Exec(
		queries.UpdateItem,
		item.Item_Name,
		item.Unit_Price,
		item.Units,
		item.Store_Branch,
		item.Weight,
		item.Category,
		item.Subcategory,
		item.Is_Organic,
		item.Brand_Name,
		item.Barcode,
		item.Updated_At,
		item.Item_ID,
	)
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
			&item.Category,
			&item.Subcategory,
			&item.Is_Organic,
			&item.Brand_Name,
			&item.Barcode,
			&item.Updated_At,
		)
		if err != nil {
			log.Printf("Scan error: %v", err)
			return nil, fmt.Errorf("error scanning item: %w", err)
		}
		items = append(items, &item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating through items: %w", err)
	}

	if len(items) == 0 {
		return []*models.Item{}, nil
	}

	return items, nil
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
