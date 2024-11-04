package models

import (
	"time"
)

// Item represents an item in the inventory.
type Item struct {
	ItemID    int       `json:"item_id"`    // Unique identifier for each item
	ItemName  string    `json:"item_name"`  // Name or description of the item
	UnitPrice float64   `json:"unit_price"` // Price per unit of the item
	Units     float64   `json:"units"`      // Number of units available
	VendorID  int       `json:"vendor_id"`  // Foreign key linking to the Vendors table
	Weight    float64   `json:"weight"`     // Weight of the item
	CreatedAt time.Time `json:"created_at"` // Timestamp for when the item was created
	UpdatedAt time.Time `json:"updated_at"` // Timestamp for last update
}
