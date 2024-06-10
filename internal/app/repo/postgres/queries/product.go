package queries

const (
	QueryCreateProduct = `
		INSERT INTO products (name, category, price)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	QueryGetProductByID = `
		SELECT id, name, category_id, price, created_at, updated_at
		FROM products
		WHERE id = $1
	`

	QueryGetPriceByProductID = `
		SELECT price
		FROM products
		WHERE id = $1
	`

	QueryGetProductByCategoryID = `
		SELECT id, name, category_id, price, created_at, updated_at
		FROM products
		WHERE category_id = $1
	`

	QueryGetAllProducts = `
		SELECT COUNT(*) OVER(), id, name, category_id, price, created_at, updated_at
		FROM products
		LIMIT $1 OFFSET $2
	`

	QueryUpdateProductPrice = `
		UPDATE products
		SET price = $1, updated_at = NOW()
		WHERE id = $2
		RETURNING id
	`

	QueryDeleteProduct = `
		UPDATE products
		SET deleted_at = NOW()
		WHERE id = $1
	`
)
