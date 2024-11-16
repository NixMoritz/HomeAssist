package models

import "time"

type Store struct {
	StoreID      int       `json:"store_id"`      // Unique identifier for each store
	UniqueUID    string    `json:"unique_uid"`    // Unique identifier
	StoreBranch  string    `json:"store_branch"`  // Chain name (e.g., REWE, Kaufland)
	StoreName    string    `json:"store_name"`    // Optional name (e.g., Edeka Special)
	StoreAddress string    `json:"store_address"` // Address of the store
	StorePhone   string    `json:"store_phone"`   // Contact phone number for the store
	CreatedAt    time.Time `json:"created_at"`    // Timestamp for when the store was created
	UpdatedAt    time.Time `json:"updated_at"`    // Timestamp for last update
}
type Item struct {
	ItemID      int       `json:"item_id"`      // Unique identifier for each item
	ItemName    string    `json:"item_name"`    // Name or description of the item
	UnitPrice   float64   `json:"unit_price"`   // Price per unit of the item
	Units       float64   `json:"units"`        // Units of the item (default 1)
	StoreBranch string    `json:"store_branch"` // Chain name (e.g., REWE, Kaufland)
	Weight      float64   `json:"weight"`       // Weight of the item
	CreatedAt   time.Time `json:"created_at"`   // Timestamp for when the item was created
	UpdatedAt   time.Time `json:"updated_at"`   // Timestamp for last update
}

type Receipt struct {
	ReceiptID           int       `json:"receipt_id"`            // Unique identifier for each receipt
	StoreID             int       `json:"store_id"`              // Foreign key linking to the Stores table
	DateIssued          time.Time `json:"date_issued"`           // Date when the receipt was issued
	TotalAmount         float64   `json:"total_amount"`          // Total amount of the receipt
	PaymentMethod       string    `json:"payment_method"`        // Payment method used (e.g., Cash, Credit Card)
	TotalDiscountAmount float64   `json:"total_discount_amount"` // Discount amount applied, if any
	NetAmount           float64   `json:"net_amount"`            // Final amount after applying tax and discount
	Notes               string    `json:"notes"`                 // Additional notes or comments
	CreatedAt           time.Time `json:"created_at"`            // Timestamp for when the receipt was created
	UpdatedAt           time.Time `json:"updated_at"`            // Timestamp for last update
}

type ReceiptItem struct {
	ReceiptItemID  int     `json:"receipt_item_id"` // Unique identifier for each receipt-item entry
	ReceiptID      int     `json:"receipt_id"`      // Foreign key linking to the Receipts table
	ItemID         int     `json:"item_id"`         // Foreign key linking to the Items table
	Quantity       int     `json:"quantity"`        // Quantity of the item purchased
	DiscountAmount float64 `json:"discount_amount"` // Discount amount applied to the item
	TotalPrice     float64 `json:"total_price"`     // Total price for the item (UnitPrice * Quantity)
}
