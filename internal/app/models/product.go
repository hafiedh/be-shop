package models

type (
	Product struct {
		ID         int     `json:"id,omitempty"`
		Name       string  `json:"name" validate:"required"`
		CategoryID string  `json:"category_id" validate:"required"`
		Price      float64 `json:"price" validate:"required,gt=0"`
		CreatedAt  string  `json:"created_at,omitempty"`
		UpdatedAt  string  `json:"updated_at,omitempty"`
	}
)
