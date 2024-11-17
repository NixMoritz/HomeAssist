package models

import "time"

type Store struct {
	Store_ID      int       `json:"store_id"`      // Unique identifier for each store
	Unique_UID    string    `json:"unique_uid"`    // Unique identifier
	Store_Branch  string    `json:"store_branch"`  // Chain name (e.g., REWE, Kaufland)
	Store_Name    string    `json:"store_name"`    // Optional name (e.g., Edeka Special)
	Store_Address string    `json:"store_address"` // Address of the store
	Store_Phone   string    `json:"store_phone"`   // Contact phone number for the store
	Updated_At    time.Time `json:"updated_at"`    // Timestamp for last update
}
type Item struct {
	Item_ID      int       `json:"item_id"`      // Unique identifier for each item
	Item_Name    string    `json:"item_name"`    // Name or description of the item
	Unit_Price   float64   `json:"unit_price"`   // Price per unit of the item
	Units        float64   `json:"units"`        // Units of the item (default 1)
	Store_Branch string    `json:"store_branch"` // Chain name (e.g., REWE, Kaufland)
	Weight       float64   `json:"weight"`       // Weight of the item
	Updated_At   time.Time `json:"updated_at"`   // Timestamp for last update
}

type Receipt struct {
	Receipt_ID            int       `json:"receipt_id"`            // Unique identifier for each receipt
	Store_ID              int       `json:"store_id"`              // Foreign key linking to the Stores table
	Date_Issued           time.Time `json:"date_issued"`           // Date when the receipt was issued
	Total_Amount          float64   `json:"total_amount"`          // Total amount of the receipt
	Payment_Method        string    `json:"payment_method"`        // Payment method used (e.g., Cash, Credit Card)
	Total_Discount_Amount float64   `json:"total_discount_amount"` // Discount amount applied, if any
	Net_Amount            float64   `json:"net_amount"`            // Final amount after applying tax and discount
	Notes                 string    `json:"notes"`                 // Additional notes or comments
	Updated_At            time.Time `json:"updated_at"`            // Timestamp for last update
}

type ReceiptItem struct {
	Receipt_Item_ID int     `json:"receipt_item_id"` // Unique identifier for each receipt-item entry
	Receipt_ID      int     `json:"receipt_id"`      // Foreign key linking to the Receipts table
	Item_ID         int     `json:"item_id"`         // Foreign key linking to the Items table
	Quantity        int     `json:"quantity"`        // Quantity of the item purchased
	Discount_Amount float64 `json:"discount_amount"` // Discount amount applied to the item
	Total_Price     float64 `json:"total_price"`     // Total price for the item (UnitPrice * Quantity)
}
