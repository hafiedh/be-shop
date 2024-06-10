package controller

import (
	"be-shop/internal/app/models"
	"be-shop/internal/app/service"
	"be-shop/internal/app/service/utils"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type (
	PaymentCtrl interface {
		Checkout(ec echo.Context) error
		SimulatePayment(ec echo.Context) error
	}

	PaymentCtrlImpl struct {
		dig.In

		PaymentSvc service.PaymentSvc
	}
)

func NewPaymentCtrl(impl PaymentCtrlImpl) PaymentCtrl {
	return &impl
}

func (p *PaymentCtrlImpl) Checkout(ec echo.Context) error {
	Recover()
	ctx := ec.Request().Context()

	resp, err := p.PaymentSvc.CreatePayment(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "[PaymentCtrl.Checkout] error while CreatePayment err", "%v", err.Error())
		return ec.JSON(resp.Code, resp)
	}

	return ec.JSON(resp.Code, resp)
}

func (p *PaymentCtrlImpl) SimulatePayment(ec echo.Context) error {
	Recover()
	ctx := ec.Request().Context()

	var req service.SimulationPaymentReq

	if err := ec.Bind(&req); err != nil {
		slog.Error("SimulatePayment - Invalid request body", err)
		return ec.JSON(http.StatusBadRequest, models.DefaultResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	validate := utils.Validate

	err := validate.Struct(req)
	if err != nil {
		slog.Error("SimulatePayment - validation error", err)
		errors := err.(validator.ValidationErrors)
		return ec.JSON(http.StatusBadRequest, models.DefaultResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
			Error:   errors.Error(),
		})
	}

	resp, err := p.PaymentSvc.SimulationPayment(ctx, req)
	if err != nil {
		slog.ErrorContext(ctx, "[PaymentCtrl.SimulatePayment] error while SimulationPayment err", "%v", err.Error())
		return ec.JSON(resp.Code, resp)
	}

	return ec.JSON(resp.Code, resp)
}
