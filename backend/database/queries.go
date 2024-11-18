package database

const (
	InsertItemQuery = `
		INSERT INTO items (item_name, unit_price, units, store_branch, weight)
		VALUES ($1, $2, $3, $4, $5)
	`
	InsertStoreQuery = `
		INSERT INTO stores (unique_uid, store_branch, store_name, store_address, store_phone)
		VALUES ($1, $2, $3, $4, $5)
	`

	InsertReceiptQuery = `
		INSERT INTO receipts (receipt_id, store_id, date_issued, total_amount, payment_method, total_discount_amount, net_amount, notes )
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	InsertReceiptItemQuery = `
		INSERT INTO receipts_items (receipt_item_id, receipt_id, item_id, quantity, discount_amount, total_price)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	getItem = `
		SELECT * FROM items
		WHERE item_id = $1
		`

	getAllItems = `
		SELECT * FROM items
		`

	getStore = `
		SELECT * FROM stores
		WHERE store_id = $1
		`

	getAllStores = `
		SELECT * FROM stores
		`

	getReceipt = `
		SELECT * FROM receipts
		WHERE receipt_id = $1
		`
)
