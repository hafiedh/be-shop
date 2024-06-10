package service

import (
	"be-shop/internal/app/models"
	"be-shop/internal/app/repo/postgres"
	"be-shop/pkg/middleware"
	"context"
	"log/slog"
	"net/http"

	"go.uber.org/dig"
)

type (
	UpdateCartQuantityReq struct {
		Quantity int `json:"quantity" validate:"required,gt=0"`
	}

	AddToCartReq struct {
		ProductID int `json:"product_id" validate:"required"`
		Quantity  int `json:"quantity" validate:"required,gt=0"`
	}
	CartSvc interface {
		AddToCart(ctx context.Context, req AddToCartReq) (resp models.DefaultResponse, err error)
		GetCart(ctx context.Context) (resp models.DefaultResponse, err error)
		UpdateCartQuantity(ctx context.Context, id int64, req UpdateCartQuantityReq) (resp models.DefaultResponse, err error)
		DeleteCart(ctx context.Context, id int64) (resp models.DefaultResponse, err error)
		DeleteAllCart(ctx context.Context) (resp models.DefaultResponse, err error)
	}

	CartSvcImpl struct {
		dig.In

		CartRepo    postgres.CartRepo
		ProductRepo postgres.ProductRepo
	}
)

func NewCardSvc(impl CartSvcImpl) CartSvc {
	return &impl
}

func (c *CartSvcImpl) AddToCart(ctx context.Context, req AddToCartReq) (resp models.DefaultResponse, err error) {
	{
		resp.Message = "Failed to add product to cart"
		resp.Code = http.StatusBadGateway
	}
	userData, ok := ctx.Value(middleware.UserData).(middleware.UserCtxReq)
	if !ok {
		slog.ErrorContext(ctx, "[CartSvcImpl.AddToCart] error while get user data")
		resp.Message = "Failed to add product to cart"
		resp.Code = http.StatusUnauthorized
		return
	}

	_, err = c.ProductRepo.GetProductByID(ctx, int64(req.ProductID))
	if err != nil {
		slog.ErrorContext(ctx, "[CartSvcImpl.AddToCart] error while GetProductByID err", "%v", err.Error())
		resp.Message = "Product not found"
		resp.Code = http.StatusNotFound
		return
	}

	entryCart := models.Cart{
		UserID:    int(userData.UserID),
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}

	id, err := c.CartRepo.CreateCart(ctx, entryCart)
	if err != nil {
		slog.ErrorContext(ctx, "[CartSvcImpl.AddToCart] error while CreateCart err", "%v", err.Error())
		resp.Message = "Failed to add product to cart"
		resp.Code = http.StatusBadGateway
		return
	}

	resp.Message = "Product added to cart successfully"
	resp.Code = http.StatusCreated
	resp.Data = struct {
		ID int `json:"id"`
	}{
		ID: id,
	}
	return
}

func (c *CartSvcImpl) GetCart(ctx context.Context) (resp models.DefaultResponse, err error) {
	{
		resp.Message = "Failed to get cart"
		resp.Code = http.StatusBadGateway
	}

	userData, ok := ctx.Value(middleware.UserData).(middleware.UserCtxReq)
	if !ok {
		slog.ErrorContext(ctx, "[CartSvcImpl.GetCart] error while get user data")
		resp.Message = "Failed to get cart"
		resp.Code = http.StatusUnauthorized
		return
	}

	carts, err := c.CartRepo.GetCartByUserID(ctx, int64(userData.UserID))
	if err != nil {
		resp.Message = "Failed to get cart"
		resp.Code = http.StatusBadGateway
		return
	}

	// should be cached here to avoid multiple request to database

	resp.Message = "Success"
	resp.Code = http.StatusOK
	resp.Data = carts
	return
}

func (c *CartSvcImpl) UpdateCartQuantity(ctx context.Context, id int64, req UpdateCartQuantityReq) (resp models.DefaultResponse, err error) {
	{
		resp.Message = "Failed to update cart"
		resp.Code = http.StatusBadGateway
	}

	userData, ok := ctx.Value(middleware.UserData).(middleware.UserCtxReq)
	if !ok {
		slog.ErrorContext(ctx, "[CartSvcImpl.UpdateCartQuantity] error while get user data")
		resp.Message = "Failed to update cart"
		resp.Code = http.StatusUnauthorized
		return
	}

	err = c.CartRepo.UpdateCartQuantity(ctx, id, int64(userData.UserID), req.Quantity)
	if err != nil {
		slog.ErrorContext(ctx, "[CartSvcImpl.UpdateCartQuantity] error while UpdateCartQuantity err", "%v", err.Error())
		resp.Message = "Failed to update cart"
		resp.Code = http.StatusBadGateway
		return
	}

	// should be removed from cache here to make sure the data is up to date

	resp.Message = "Cart updated successfully"
	resp.Code = http.StatusOK
	return
}

func (c *CartSvcImpl) DeleteCart(ctx context.Context, id int64) (resp models.DefaultResponse, err error) {
	{
		resp.Message = "Failed to delete cart"
		resp.Code = http.StatusBadGateway
	}

	userData, ok := ctx.Value(middleware.UserData).(middleware.UserCtxReq)
	if !ok {
		slog.ErrorContext(ctx, "[CartSvcImpl.DeleteCart] error while get user data")
		resp.Message = "Failed to delete cart"
		resp.Code = http.StatusUnauthorized
		return
	}

	err = c.CartRepo.DeleteCart(ctx, int64(userData.UserID), id)
	if err != nil {
		slog.ErrorContext(ctx, "[CartSvcImpl.DeleteCart] error while DeleteCart err", "%v", err.Error())
		resp.Message = "Failed to delete cart"
		resp.Code = http.StatusBadGateway
		return
	}

	// should be removed from cache here to make sure the data is up to date

	resp.Message = "Cart deleted successfully"
	resp.Code = http.StatusOK
	return
}

func (c *CartSvcImpl) DeleteAllCart(ctx context.Context) (resp models.DefaultResponse, err error) {
	{
		resp.Message = "Failed to delete cart"
		resp.Code = http.StatusBadGateway
	}

	userData, ok := ctx.Value(middleware.UserData).(middleware.UserCtxReq)
	if !ok {
		slog.ErrorContext(ctx, "[CartSvcImpl.DeleteAllCart] error while get user data")
		resp.Message = "Failed to delete cart"
		resp.Code = http.StatusUnauthorized
		return
	}

	err = c.CartRepo.DeleteAllCart(ctx, int64(userData.UserID))
	if err != nil {
		slog.ErrorContext(ctx, "[CartSvcImpl.DeleteAllCart] error while DeleteAllCart err", "%v", err.Error())
		resp.Message = "Failed to delete cart"
		resp.Code = http.StatusBadGateway
		return
	}

	// should be removed from cache here to make sure the data is up to date

	resp.Message = "Cart deleted successfully"
	resp.Code = http.StatusOK
	return
}
