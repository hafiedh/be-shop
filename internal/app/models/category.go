package models

type (
	Category struct {
		ID        int    `json:"id,omitempty"`
		Name      string `json:"name" validate:"required"`
		CreatedAt string `json:"created_at,omitempty"`
		UpdatedAt string `json:"updated_at,omitempty"`
	}
)
