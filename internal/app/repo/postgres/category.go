package postgres

import (
	"be-shop/internal/app/repo/postgres/queries"
	"context"
	"database/sql"

	"go.uber.org/dig"
)

type (
	CategoryRepo interface {
		CreateCategory(ctx context.Context, name string) (id int, err error)
	}

	CategoryRepoImpl struct {
		dig.In

		*sql.DB
	}
)

func NewCategoryRepo(impl CategoryRepoImpl) CategoryRepo {
	return &impl
}

func (c *CategoryRepoImpl) CreateCategory(ctx context.Context, name string) (id int, err error) {
	err = c.QueryRowContext(ctx, queries.QueryCreateCategory, name).Scan(&id)
	return
}
