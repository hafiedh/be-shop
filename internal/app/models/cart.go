package models

type (
	Cart struct {
		ID           int     `json:"id,omitempty"`
		UserID       int     `json:"user_id" validate:"required"`
		ProductID    int     `json:"product_id" validate:"required"`
		ProductName  string  `json:"product_name"`
		ProductPrice float64 `json:"product_price,omitempty"`
		Quantity     int     `json:"quantity" validate:"required"`
		CreatedAt    string  `json:"created_at,omitempty"`
		UpdatedAt    string  `json:"updated_at,omitempty"`
	}
)
