package controller

import (
	"be-shop/internal/app/models"
	"be-shop/internal/app/service"
	"be-shop/internal/app/service/utils"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type (
	CartCtrl interface {
		AddToCart(ec echo.Context) error
		GetCart(ec echo.Context) error
		UpdateCartQuantity(ec echo.Context) error
		DeleteCart(ec echo.Context) error
		DeleteAllCart(ec echo.Context) error
	}

	CartCtrlImpl struct {
		dig.In

		CartSvc service.CartSvc
	}
)

func NewCartCtrl(impl CartCtrlImpl) CartCtrl {
	return &impl
}

func (m *CartCtrlImpl) AddToCart(ec echo.Context) error {
	Recover()
	ctx := ec.Request().Context()

	var req service.AddToCartReq
	if err := ec.Bind(&req); err != nil {
		slog.Error("AddToCart - Invalid request body", err)
		return ec.JSON(http.StatusBadRequest, models.DefaultResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	validate := utils.Validate

	err := validate.Struct(req)
	if err != nil {
		slog.Error("AddToCart - validation error", err)
		errors := err.(validator.ValidationErrors)
		return ec.JSON(http.StatusBadRequest, models.DefaultResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
			Error:   errors.Error(),
		})
	}

	resp, err := m.CartSvc.AddToCart(ctx, req)
	if err != nil {
		slog.Error("AddToCart - error", err)
		return ec.JSON(resp.Code, resp)
	}

	return ec.JSON(resp.Code, resp)
}

func (m *CartCtrlImpl) GetCart(ec echo.Context) error {
	Recover()

	ctx := ec.Request().Context()

	resp, err := m.CartSvc.GetCart(ctx)
	if err != nil {
		slog.Error("GetCart - error", err)
		return ec.JSON(resp.Code, resp)
	}

	return ec.JSON(resp.Code, resp)
}

func (m *CartCtrlImpl) UpdateCartQuantity(ec echo.Context) error {
	Recover()

	ctx := ec.Request().Context()

	id := ec.Param("id")

	idConv, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		slog.Error("UpdateCartQuantity - error while converting id", err)
		return ec.JSON(http.StatusBadRequest, models.DefaultResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	var req service.UpdateCartQuantityReq

	if err := ec.Bind(&req); err != nil {
		slog.Error("UpdateCartQuantity - Invalid request body", err)
		return ec.JSON(http.StatusBadRequest, models.DefaultResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	validate := utils.Validate

	err = validate.Struct(req)
	if err != nil {
		slog.Error("UpdateCartQuantity - validation error", err)
		errors := err.(validator.ValidationErrors)
		return ec.JSON(http.StatusBadRequest, models.DefaultResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
			Error:   errors.Error(),
		})
	}

	resp, err := m.CartSvc.UpdateCartQuantity(ctx, idConv, req)
	if err != nil {
		slog.Error("UpdateCartQuantity - error", err)
		return ec.JSON(resp.Code, resp)
	}

	return ec.JSON(resp.Code, resp)

}

func (m *CartCtrlImpl) DeleteCart(ec echo.Context) error {
	Recover()

	ctx := ec.Request().Context()

	id := ec.Param("id")

	idConv, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		slog.Error("DeleteCart - error while converting id", err)
		return ec.JSON(http.StatusBadRequest, models.DefaultResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	resp, err := m.CartSvc.DeleteCart(ctx, idConv)
	if err != nil {
		slog.Error("DeleteCart - error", err)
		return ec.JSON(resp.Code, resp)
	}

	return ec.JSON(resp.Code, resp)
}

func (m *CartCtrlImpl) DeleteAllCart(ec echo.Context) error {
	Recover()

	ctx := ec.Request().Context()

	resp, err := m.CartSvc.DeleteAllCart(ctx)
	if err != nil {
		slog.Error("DeleteAllCart - error", err)
		return ec.JSON(resp.Code, resp)
	}

	return ec.JSON(resp.Code, resp)
}
