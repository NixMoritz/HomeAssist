package database

import (
	"HomeAssist/internal/models"
	"HomeAssist/internal/storage/database/queries"
	"database/sql"
	"fmt"
)

func AddNewReceipt(receipt models.Receipt, db *sql.DB) error {
	err := db.QueryRow(
		queries.InsertReceiptQuery,
		receipt.Store_ID,
		receipt.Date_Issued,
		receipt.Total_Amount,
		receipt.Payment_Method,
		receipt.Total_Discount_Amount,
		receipt.Net_Amount,
		receipt.Tax_Amount,
		receipt.Receipt_Type,
		receipt.Loyalty_Card_Number,
		receipt.Cashier_Name,
		receipt.Receipt_Number,
		receipt.Notes,
		receipt.Image_URL,
	).Scan(&receipt.Receipt_ID)

	if err != nil {
		return fmt.Errorf("failed to insert receipt: %v", err)
	}

	return nil
}

func GetReceipt(receiptID int, db *sql.DB) (models.Receipt, error) {
	var receipt models.Receipt
	err := db.QueryRow(queries.GetReceiptQuery, receiptID).Scan(
		&receipt.Receipt_ID,
		&receipt.Store_ID,
		&receipt.Date_Issued,
		&receipt.Total_Amount,
		&receipt.Payment_Method,
		&receipt.Total_Discount_Amount,
		&receipt.Net_Amount,
		&receipt.Tax_Amount,
		&receipt.Receipt_Type,
		&receipt.Loyalty_Card_Number,
		&receipt.Cashier_Name,
		&receipt.Receipt_Number,
		&receipt.Notes,
		&receipt.Image_URL,
		&receipt.Updated_At,
	)

	if err == sql.ErrNoRows {
		return receipt, fmt.Errorf("receipt with ID %d not found", receiptID)
	}
	if err != nil {
		return receipt, fmt.Errorf("error fetching receipt: %v", err)
	}

	return receipt, nil
}

func GetAllReceipts(db *sql.DB) ([]models.Receipt, error) {
	rows, err := db.Query(queries.GetAllReceiptsQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying receipts: %v", err)
	}
	defer rows.Close()

	var receipts []models.Receipt
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
			&receipt.Tax_Amount,
			&receipt.Receipt_Type,
			&receipt.Loyalty_Card_Number,
			&receipt.Cashier_Name,
			&receipt.Receipt_Number,
			&receipt.Notes,
			&receipt.Image_URL,
			&receipt.Updated_At,
			&receipt.Store_Name,
			&receipt.Store_Branch,
			&receipt.Store_Address,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning receipt: %v", err)
		}
		receipts = append(receipts, receipt)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating receipts: %v", err)
	}

	return receipts, nil
}

func UpdateReceipt(receipt models.Receipt, db *sql.DB) error {
	result, err := db.Exec(
		queries.UpdateReceiptQuery,
		receipt.Store_ID,
		receipt.Date_Issued,
		receipt.Total_Amount,
		receipt.Payment_Method,
		receipt.Total_Discount_Amount,
		receipt.Net_Amount,
		receipt.Tax_Amount,
		receipt.Receipt_Type,
		receipt.Loyalty_Card_Number,
		receipt.Cashier_Name,
		receipt.Receipt_Number,
		receipt.Notes,
		receipt.Image_URL,
		receipt.Receipt_ID,
	)

	if err != nil {
		return fmt.Errorf("failed to update receipt: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking update result: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("receipt with ID %d not found", receipt.Receipt_ID)
	}

	return nil
}

func DeleteReceipt(receiptID int, db *sql.DB) error {
	// First delete related receipt items
	_, err := db.Exec("DELETE FROM RECEIPT_ITEMS WHERE RECEIPT_ID = $1", receiptID)
	if err != nil {
		return fmt.Errorf("failed to delete receipt items: %v", err)
	}

	// Then delete the receipt
	result, err := db.Exec(queries.DeleteReceiptQuery, receiptID)
	if err != nil {
		return fmt.Errorf("failed to delete receipt: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking delete result: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("receipt with ID %d not found", receiptID)
	}

	return nil
}
