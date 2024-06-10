package models

type (
	Order struct {
		ID          int     `json:"id,omitempty"`
		UserID      int     `json:"user_id" validate:"required"`
		TotalAmount float64 `json:"total_amount" validate:"required"`
		OrderCode   string  `json:"order_code" validate:"required"`
		Status      string  `json:"status"`
		CreatedAt   string  `json:"created_at,omitempty"`
		UpdatedAt   string  `json:"updated_at,omitempty"`
	}
)
