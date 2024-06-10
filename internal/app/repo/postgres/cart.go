package postgres

import (
	"be-shop/internal/app/models"
	"be-shop/internal/app/repo/postgres/queries"
	"context"
	"database/sql"
	"log/slog"

	"go.uber.org/dig"
)

type (
	CartRepo interface {
		CreateCart(ctx context.Context, req models.Cart) (id int, err error)
		GetCartByUserID(ctx context.Context, userID int64) (resp []models.Cart, err error)
		UpdateCartQuantity(ctx context.Context, userID, id int64, quantity int) (err error)
		DeleteCart(ctx context.Context, userID, id int64) (err error)
		DeleteAllCart(ctx context.Context, userID int64) (err error)
	}

	CartRepoImpl struct {
		dig.In

		*sql.DB
	}
)

func NewCartRepo(impl CartRepoImpl) CartRepo {
	return &impl
}

func (c *CartRepoImpl) CreateCart(ctx context.Context, req models.Cart) (id int, err error) {
	row := c.QueryRowContext(ctx, queries.QueryCreateCart, req.UserID, req.ProductID, req.Quantity)
	err = row.Scan(&id)
	if err != nil {
		slog.ErrorContext(ctx, "[CartRepoImpl.CreateCart] error while CreateCart err", "%v", err.Error())
		return
	}
	return
}

func (c *CartRepoImpl) GetCartByUserID(ctx context.Context, userID int64) (resp []models.Cart, err error) {
	rows, err := c.QueryContext(ctx, queries.QueryGetCartByUserID, userID)
	if err != nil {
		slog.ErrorContext(ctx, "[CartRepoImpl.GetCartByUserID] error while GetCartByUserID err", "%v", err.Error())
		return
	}

	for rows.Next() {
		var cart models.Cart
		err = rows.Scan(&cart.ID, &cart.UserID, &cart.ProductID, &cart.ProductName, &cart.Quantity)
		if err != nil {
			slog.ErrorContext(ctx, "[CartRepoImpl.GetCartByUserID] error while scan err", "%v", err.Error())
			return
		}
		resp = append(resp, cart)
	}

	return
}

func (c *CartRepoImpl) UpdateCartQuantity(ctx context.Context, userID, id int64, quantity int) (err error) {
	_, err = c.ExecContext(ctx, queries.QueryUpdateCartQuantity, quantity, id, userID)
	if err != nil {
		slog.ErrorContext(ctx, "[CartRepoImpl.UpdateCartQuantity] error while UpdateCartQuantity err", "%v", err.Error())
		return
	}
	return
}

func (c *CartRepoImpl) DeleteCart(ctx context.Context, userID, id int64) (err error) {
	_, err = c.ExecContext(ctx, queries.QueryDeleteCart, id, userID)
	if err != nil {
		slog.ErrorContext(ctx, "[CartRepoImpl.DeleteCart] error while DeleteCart err", "%v", err.Error())
		return
	}
	return
}

func (c *CartRepoImpl) DeleteAllCart(ctx context.Context, userID int64) (err error) {
	_, err = c.ExecContext(ctx, queries.QueryDeleteAllCart, userID)
	if err != nil {
		slog.ErrorContext(ctx, "[CartRepoImpl.DeleteAllCart] error while DeleteAllCart err", "%v", err.Error())
		return
	}
	return
}
