package queries

const (
	QueryCreateCart = `INSERT INTO cart_items (user_id, product_id, quantity) VALUES ($1, $2, $3) 
	ON CONFLICT (user_id, product_id) DO UPDATE SET quantity = cart_items.quantity + $3 RETURNING id`
	QueryGetCartByUserID = `SELECT c.id, c.user_id, c.product_id, p.name, c.quantity FROM cart_items c
	JOIN products p ON c.product_id = p.id WHERE c.user_id = $1`

	QueryUpdateCartQuantity = `UPDATE cart_items SET quantity = $1 WHERE id = $2 AND user_id = $3`

	QueryDeleteCart = `DELETE FROM cart_items WHERE id = $1 AND user_id = $2`

	QueryDeleteAllCart = `DELETE FROM cart_items WHERE user_id = $1`
)
