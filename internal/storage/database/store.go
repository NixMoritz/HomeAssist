package database

import (
	"database/sql"
	"fmt"
	"log"

	"HomeAssist/internal/models"
	"HomeAssist/internal/storage/database/queries"
)

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

func GetStoreByID(storeID int, db *sql.DB) (models.Store, error) {
	var store models.Store
	err := db.QueryRow(`
		SELECT store_id, store_name, store_branch, store_address, store_phone, 
			   store_type, store_category, updated_at 
		FROM stores 
		WHERE store_id = $1`, storeID).Scan(
		&store.Store_ID,
		&store.Store_Name,
		&store.Store_Branch,
		&store.Store_Address,
		&store.Store_Phone,
		&store.Store_Type,
		&store.Store_Category,
		&store.Updated_At,
	)
	if err == sql.ErrNoRows {
		return store, fmt.Errorf("store with ID %d not found", storeID)
	}
	if err != nil {
		return store, fmt.Errorf("error fetching store: %v", err)
	}
	return store, nil
}

func GetAllStores(db *sql.DB) ([]models.Store, error) {
	rows, err := db.Query(queries.GetAllStoresQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying stores: %v", err)
	}
	defer rows.Close()

	var stores []models.Store
	for rows.Next() {
		var store models.Store
		var storeType, storeCategory sql.NullString
		var uniqueUID sql.NullString

		err := rows.Scan(
			&store.Store_ID,
			&uniqueUID,
			&store.Store_Name,
			&store.Store_Branch,
			&store.Store_Address,
			&store.Store_Phone,
			&storeType,
			&storeCategory,
			&store.Updated_At,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning store: %v", err)
		}

		// Convert NullString to string, using empty string if null
		store.Store_Type = storeType.String
		store.Store_Category = storeCategory.String
		store.Unique_UID = uniqueUID.String

		stores = append(stores, store)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating stores: %v", err)
	}

	return stores, nil
}
