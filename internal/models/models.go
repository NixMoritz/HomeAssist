package models

import (
	"time"
)

type Store struct {
	Store_ID       int       `json:"store_id"`                 // Unique identifier for each store
	Unique_UID     string    `json:"unique_uid,omitempty"`     // Unique identifier
	Store_Branch   string    `json:"store_branch"`             // Chain name (e.g., REWE, Kaufland)
	Store_Name     string    `json:"store_name"`               // Optional name (e.g., Edeka Special)
	Store_Address  string    `json:"store_address"`            // Address of the store
	Store_Phone    string    `json:"store_phone"`              // Contact phone number for the store
	Store_Type     string    `json:"store_type,omitempty"`     // Make optional
	Store_Category string    `json:"store_category,omitempty"` // Make optional
	Updated_At     time.Time `json:"updated_at"`               // Timestamp for last update
}

type Item struct {
	Item_ID      int       `json:"item_id"`
	Item_Name    string    `json:"item_name"`
	Unit_Price   float64   `json:"unit_price"`
	Units        float64   `json:"units"`
	Store_Branch string    `json:"store_branch"`
	Weight       float64   `json:"weight"`
	Category     *string   `json:"category"`
	Subcategory  *string   `json:"subcategory"`
	Is_Organic   *bool     `json:"is_organic"`
	Brand_Name   *string   `json:"brand_name"`
	Barcode      *string   `json:"barcode"`
	Updated_At   time.Time `json:"updated_at"`
}

type Receipt struct {
	Receipt_ID            int       `json:"receipt_id"`
	Store_ID              int       `json:"store_id"`
	Store_Name            string    `json:"store_name"`
	Store_Branch          string    `json:"store_branch"`
	Store_Address         string    `json:"store_address"`
	Date_Issued           time.Time `json:"date_issued"` // This will store both date and time
	Total_Amount          float64   `json:"total_amount"`
	Payment_Method        string    `json:"payment_method"`
	Total_Discount_Amount float64   `json:"total_discount_amount"`
	Net_Amount            float64   `json:"net_amount"`
	Tax_Amount            float64   `json:"tax_amount"`
	Receipt_Type          string    `json:"receipt_type"`
	Loyalty_Card_Number   string    `json:"loyalty_card_number"`
	Cashier_Name          string    `json:"cashier_name"`
	Receipt_Number        string    `json:"receipt_number"`
	Notes                 string    `json:"notes"`
	Image_URL             string    `json:"image_url"`
	Updated_At            time.Time `json:"updated_at"`
}

type ReceiptItem struct {
	Receipt_Item_ID     int     `json:"receipt_item_id"` // Unique identifier for each receipt-item entry
	Receipt_ID          int     `json:"receipt_id"`      // Foreign key linking to the Receipts table
	Item_ID             int     `json:"item_id"`         // Foreign key linking to the Items table
	Quantity            int     `json:"quantity"`        // Quantity of the item purchased
	Unit_Price          float64 `json:"unit_price"`      // Added
	Discount_Amount     float64 `json:"discount_amount"`
	Discount_Percentage float64 `json:"discount_percentage"` // Added
	Tax_Rate            float64 `json:"tax_rate"`            // Added
	Tax_Amount          float64 `json:"tax_amount"`          // Added
	Total_Price         float64 `json:"total_price"`
	Weight_At_Purchase  float64 `json:"weight_at_purchase"` // Added
	Notes               string  `json:"notes"`              // Added
}

// Added: New struct for promotions
type Promotion struct {
	Promotion_ID          int       `json:"promotion_id"`
	Receipt_Item_ID       int       `json:"receipt_item_id"`
	Promotion_Type        string    `json:"promotion_type"`
	Promotion_Description string    `json:"promotion_description"`
	Discount_Amount       float64   `json:"discount_amount"`
	Valid_From            time.Time `json:"valid_from"`
	Valid_To              time.Time `json:"valid_to"`
}

type ReceiptEntry struct {
	Receipt      Receipt       `json:"receipt"`
	ReceiptItems []ReceiptItem `json:"receipt_items"`
	Promotions   []Promotion   `json:"promotions,omitempty"` // Added: Optional promotions
}

type Items struct {
	Items []Item `json:"items"`
}
