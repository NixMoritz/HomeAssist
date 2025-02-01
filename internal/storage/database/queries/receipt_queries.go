package queries

const (
	// Get all receipts
	GetAllReceiptsQuery = `
		SELECT r.receipt_id, r.store_id, r.date_issued, r.total_amount, 
			   r.payment_method, r.total_discount_amount, r.net_amount, 
			   r.tax_amount, r.receipt_type, r.loyalty_card_number, 
			   r.cashier_name, r.receipt_number, r.notes, r.image_url, 
			   r.updated_at, s.store_name, s.store_branch, s.store_address
		FROM receipts r
		LEFT JOIN stores s ON r.store_id = s.store_id
		ORDER BY r.date_issued DESC
	`
	// Insert new receipt
	InsertReceiptQuery = `
		INSERT INTO receipts (
			store_id, date_issued, total_amount, payment_method,
			total_discount_amount, net_amount, tax_amount, receipt_type,
			loyalty_card_number, cashier_name, receipt_number, notes, image_url
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
		)
		RETURNING receipt_id
	`

	// Update existing receipt
	UpdateReceiptQuery = `
		UPDATE receipts SET 
			store_id = $1, date_issued = $2, total_amount = $3,
			payment_method = $4, total_discount_amount = $5, net_amount = $6,
			tax_amount = $7, receipt_type = $8, loyalty_card_number = $9,
			cashier_name = $10, receipt_number = $11, notes = $12,
			image_url = $13
		WHERE receipt_id = $14
		`

	// Delete receipt
	DeleteReceiptQuery = `DELETE FROM receipts WHERE receipt_id = $1`

	// Get single receipt
	GetReceiptQuery = `
		SELECT receipt_id, store_id, date_issued, total_amount,
			payment_method, total_discount_amount, net_amount, tax_amount,
			receipt_type, loyalty_card_number, cashier_name,
			receipt_number, notes, image_url, updated_at
		FROM receipts 
		WHERE receipt_id = $1
	`
)
