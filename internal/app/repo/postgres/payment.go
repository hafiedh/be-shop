package postgres

import (
	"be-shop/internal/app/models"
	"be-shop/internal/app/repo/postgres/queries"
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"math"

	"go.uber.org/dig"
)

type (
	PaymentRepo interface {
		Checkout(ctx context.Context, userID int64, orderCode string) (totalAmount float64, err error)
		GetPaymentByOrderCode(ctx context.Context, userID int64, orderCode string) (resp models.Order, err error)
		UpdatePaymentStatus(ctx context.Context, userID int64, orderCode, status string) (err error)
	}

	PaymentRepoImpl struct {
		dig.In

		*sql.DB
	}
)

func NewPaymentRepo(impl PaymentRepoImpl) PaymentRepo {
	return &impl
}

func (p *PaymentRepoImpl) Checkout(ctx context.Context, userID int64, orderCode string) (total float64, err error) {

	var (
		carts   []models.Cart
		orderID int
	)
	tx, err := p.BeginTx(ctx,
		&sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		slog.ErrorContext(ctx, "[PaymentRepoImpl.Checkout] error while begin transaction err", "%v", err.Error())
		return
	}
	defer func() {
		if err != nil {
			slog.ErrorContext(ctx, "[PaymentRepoImpl.Checkout] error occurred, rolling back transaction")
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	rows, err := tx.QueryContext(ctx, queries.QueryGetCartByUserID, userID)
	if err != nil {
		slog.ErrorContext(ctx, "[PaymentRepoImpl.Checkout] error while GetCartByUserID err", "%v", err.Error())
		return
	}

	for rows.Next() {
		var cart models.Cart
		err = rows.Scan(&cart.ID, &cart.UserID, &cart.ProductID, &cart.ProductName, &cart.Quantity)
		if err != nil {
			slog.ErrorContext(ctx, "[PaymentRepoImpl.Checkout] error while scan err", "%v", err.Error())
			return
		}
		carts = append(carts, cart)
	}

	if len(carts) == 0 {
		slog.ErrorContext(ctx, "[PaymentRepoImpl.Checkout] error while carts empty")
		err = fmt.Errorf("shopping cart is empty")
		return
	}

	for indexCart, cart := range carts {
		var price float64
		err = tx.QueryRowContext(ctx, queries.QueryGetPriceByProductID, cart.ProductID).Scan(&price)
		if err != nil {
			slog.ErrorContext(ctx, "[PaymentRepoImpl.Checkout] error while GetPriceByProductID err", "%v", err.Error())
			return
		}
		carts[indexCart].ProductPrice = price
		total += math.Ceil(price * float64(cart.Quantity))
	}

	err = tx.QueryRowContext(ctx, queries.QueryCreateOrder, userID, total, orderCode).Scan(&orderID)
	if err != nil {
		slog.ErrorContext(ctx, "[PaymentRepoImpl.Checkout] error while CreateOrder err", "%v", err.Error())
		return
	}

	for _, cart := range carts {
		_, err = tx.ExecContext(ctx, queries.QueryCreateOrderDetail, orderID, cart.ProductID, cart.Quantity, cart.ProductPrice)
		if err != nil {
			slog.ErrorContext(ctx, "[PaymentRepoImpl.Checkout] error while CreateOrderDetail err", "%v", err.Error())
			return
		}
	}

	_, err = tx.ExecContext(ctx, queries.QueryDeleteAllCart, userID)
	if err != nil {
		slog.ErrorContext(ctx, "[PaymentRepoImpl.Checkout] error while DeleteAllCart err", "%v", err.Error())
		return
	}

	return

}

func (p *PaymentRepoImpl) GetPaymentByOrderCode(ctx context.Context, userID int64, orderCode string) (resp models.Order, err error) {
	err = p.QueryRowContext(ctx, queries.QueryGetOrderByOrderCode, userID, orderCode).Scan(&resp.ID, &resp.UserID, &resp.TotalAmount, &resp.Status, &resp.OrderCode, &resp.CreatedAt, &resp.UpdatedAt)
	if err != nil {
		slog.ErrorContext(ctx, "[PaymentRepoImpl.GetPaymentByOrderCode] error while GetOrderByOrderCode err", "%v", err.Error())
		return
	}
	return
}

func (p *PaymentRepoImpl) UpdatePaymentStatus(ctx context.Context, userID int64, orderCode, status string) (err error) {
	_, err = p.ExecContext(ctx, queries.QueryUpdateOrderStatus, status, orderCode, userID)
	if err != nil {
		slog.ErrorContext(ctx, "[PaymentRepoImpl.UpdatePaymentStatus] error while UpdateOrderStatus err", "%v", err.Error())
		return
	}
	return
}
