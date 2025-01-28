package HomeAssist

const (
	InsertItemQuery = `
		INSERT INTO items (
			item_name, unit_price, units, store_branch, weight,
			category, subcategory, is_organic, brand_name, barcode, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING item_id
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
		SELECT item_id, item_name, unit_price, units, store_branch, weight, 
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
			unit_price = $2,
			units = $3,
			store_branch = $4,
			weight = $5,
			category = $6,
			subcategory = $7,
			is_organic = $8,
			brand_name = $9,
			barcode = $10,
			updated_at = $11
		WHERE item_id = $12
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
