package queries

const (
	InsertStoreQuery = `
        INSERT INTO stores (unique_uid, store_branch, store_name, store_address, store_phone)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING store_id
    `

	GetStoreQuery = `
        SELECT * FROM stores
        WHERE store_id = $1
    `

	GetAllStoresQuery = `
        SELECT * FROM stores
        ORDER BY store_name
    `
)
