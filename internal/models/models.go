package models

import (
	"time"
)

type Store struct {
	Store_ID       int       `json:"store_id"`       // Unique identifier for each store
	Unique_UID     string    `json:"unique_uid"`     // Unique identifier
	Store_Branch   string    `json:"store_branch"`   // Chain name (e.g., REWE, Kaufland)
	Store_Name     string    `json:"store_name"`     // Optional name (e.g., Edeka Special)
	Store_Address  string    `json:"store_address"`  // Address of the store
	Store_Phone    string    `json:"store_phone"`    // Contact phone number for the store
	Store_Type     string    `json:"store_type"`     // Added: e.g., 'Supermarket', 'Restaurant'
	Store_Category string    `json:"store_category"` // Added: e.g., 'Grocery', 'Fast Food'
	Updated_At     time.Time `json:"updated_at"`     // Timestamp for last update
}

type Item struct {
	Item_ID      int       `json:"item_id"`
	Item_Name    string    `json:"item_name"`
	Unit_Price   float64   `json:"unit_price"`
	Units        float64   `json:"units"`
	Store_Branch string    `json:"store_branch"`
	Weight       float64   `json:"weight"`
	Category     *string   `json:"category"`    // Using pointer for nullable string
	Subcategory  *string   `json:"subcategory"` // Using pointer for nullable string
	Is_Organic   *bool     `json:"is_organic"`  // Using pointer for nullable boolean
	Brand_Name   *string   `json:"brand_name"`  // Using pointer for nullable string
	Barcode      *string   `json:"barcode"`     // Using pointer for nullable string
	Updated_At   time.Time `json:"updated_at"`
}

type Receipt struct {
	Receipt_ID            int       `json:"receipt_id"`            // Unique identifier for each receipt
	Store_ID              int       `json:"store_id"`              // Foreign key linking to the Stores table
	Date_Issued           time.Time `json:"date_issued"`           // Date when the receipt was issued
	Time_Issued           time.Time `json:"time_issued"`           // Added
	Total_Amount          float64   `json:"total_amount"`          // Total amount of the receipt
	Payment_Method        string    `json:"payment_method"`        // Payment method used (e.g., Cash, Credit Card)
	Total_Discount_Amount float64   `json:"total_discount_amount"` // Discount amount applied, if any
	Net_Amount            float64   `json:"net_amount"`            // Final amount after applying tax and discount
	Tax_Amount            float64   `json:"tax_amount"`            // Added
	Receipt_Type          string    `json:"receipt_type"`          // Added
	Loyalty_Card_Number   string    `json:"loyalty_card_number"`   // Added
	Cashier_Name          string    `json:"cashier_name"`          // Added
	Receipt_Number        string    `json:"receipt_number"`        // Added
	Notes                 string    `json:"notes"`                 // Additional notes or comments
	Image_URL             string    `json:"image_url"`             // Added
	Updated_At            time.Time `json:"updated_at"`            // Timestamp for last update
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
