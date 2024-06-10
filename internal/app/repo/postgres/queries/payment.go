package queries

const (
	QueryCreateOrder = `
		INSERT INTO orders (user_id, total_amount,order_code)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	QueryCreateOrderDetail = `
		INSERT INTO order_items (order_id, product_id, quantity, price)
		VALUES ($1, $2, $3, $4)
	`

	QueryGetOrderByOrderCode = `
		SELECT id, user_id, total_amount, status, order_code, created_at, updated_at
		FROM orders
		WHERE user_id = $1 AND order_code = $2
	`

	QueryGetOrderDetailByOrderID = `
		SELECT id, order_id, product_id, quantity, price, created_at, updated_at
		FROM order_items
		WHERE order_id = $1
		`

	QueryUpdateOrderStatus = `
		UPDATE orders
		SET status = $1, updated_at = NOW()
		WHERE order_code = $2 AND user_id = $3
	`
)
