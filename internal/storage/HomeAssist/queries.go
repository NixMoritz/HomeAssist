package HomeAssist

const (
	InsertItemQuery = `
		INSERT INTO items (
			item_name, item_price, price_per_unit, units, store_branch, weight,
			category, subcategory, is_organic, brand_name, barcode
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`
	InsertStoreQuery = `
		INSERT INTO stores (unique_uid, store_branch, store_name, store_address, store_phone)
		VALUES ($1, $2, $3, $4, $5)
	`

	InsertReceiptQuery = `
		INSERT INTO receipts (store_id, date_issued, total_amount, payment_method, total_discount_amount, net_amount, notes )
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	InsertReceiptItemQuery = `
		INSERT INTO receipt_items (receipt_id, item_id, quantity, discount_amount, total_price)
		VALUES ($1, $2, $3, $4, $5)
	`

	GetItem = `
		SELECT * FROM items
		WHERE item_id = $1
		`

	GetAllItems = `
		SELECT item_id, item_name, item_price, price_per_unit, units, store_branch, weight, 
			   category, subcategory, is_organic, brand_name, barcode, updated_at 
		FROM items
		ORDER BY item_id
	`

	GetStore = `
		SELECT * FROM stores
		WHERE store_id = $1
		`

	GetAllStores = `
		SELECT * FROM stores
		`

	GetReceipt = `
		SELECT * FROM receipts
		WHERE receipt_id = $1
		`

	GetAllReceipts = `
		SELECT * FROM receipts
		`

	GetReceiptItem = `
		SELECT * FROM receipt_items
		WHERE Receipt_Item_ID = $1
		`

	GetAllReceiptItems = `
		SELECT * FROM receipt_items
		`

	DeleteItem = `
		DELETE FROM items 
		WHERE item_id = $1
		`

	DeleteStore = `
		DELETE FROM stores
		WHERE store_id = $1
		`

	DeleteReceipt = `
		DELETE FROM receipts 
		WHERE receipt_id = $1
		`

	DeleteReceiptItem = `
		DELETE FROM receipt_items 
		WHERE Receipt_Item_ID = $1
		`

	UpdateItem = `
		UPDATE items
		SET item_name = $1,
			item_price = $2,
			price_per_unit = $3,
			units = $4,
			store_branch = $5,
			weight = $6,
			category = $7,
			subcategory = $8,
			is_organic = $9,
			brand_name = $10,
			barcode = $11,
			updated_at = $12
		WHERE item_id = $13
	`

	UpdateStore = `
		DELETE FROM stores
		WHERE store_id = $1
		`

	UpdateReceipt = `
		DELETE FROM receipts 
		WHERE receipt_id = $1
		`

	UpdateReceiptItem = `
		DELETE FROM receipt_items 
		WHERE Receipt_Item_ID = $1
		`
)
