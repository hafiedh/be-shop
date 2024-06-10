package service

import (
	"be-shop/internal/app/models"
	"be-shop/internal/app/repo/postgres"
	"be-shop/internal/app/service/utils"
	"be-shop/pkg/middleware"
	"context"
	"log/slog"
	"net/http"
	"strings"

	"go.uber.org/dig"
)

type (
	SimulationPaymentReq struct {
		OrderCode string  `json:"order_code" validate:"required"`
		Amount    float64 `json:"amount" validate:"required"`
	}

	PaymentSvc interface {
		CreatePayment(ctx context.Context) (resp models.DefaultResponse, err error)
		SimulationPayment(ctx context.Context, req SimulationPaymentReq) (resp models.DefaultResponse, err error)
	}

	PaymentSvcImpl struct {
		dig.In

		PaymentRepo postgres.PaymentRepo
	}
)

func NewPaymentSvc(impl PaymentSvcImpl) PaymentSvc {
	return &impl
}

func (p *PaymentSvcImpl) CreatePayment(ctx context.Context) (resp models.DefaultResponse, err error) {
	{
		resp.Message = "Failed to create payment"
		resp.Code = http.StatusBadGateway
	}
	userData, ok := ctx.Value(middleware.UserData).(middleware.UserCtxReq)
	if !ok {
		slog.ErrorContext(ctx, "[PaymentSvc.CreatePayment] error while get user data")
		resp.Message = "Failed to add product to cart"
		resp.Code = http.StatusUnauthorized
		return
	}
	orderCode := utils.GenerateOrderCode(strings.Split(userData.Email, "@")[0])
	total, err := p.PaymentRepo.Checkout(ctx, int64(userData.UserID), orderCode)
	if err != nil {
		slog.ErrorContext(ctx, "[PaymentSvc.CreatePayment] error while Checkout err", "%v", err.Error())
		resp.Error = err.Error()
		return
	}

	resp.Message = "Payment created successfully"
	resp.Code = http.StatusCreated
	resp.Data = struct {
		OrderCode   string  `json:"order_code"`
		TotalAmount float64 `json:"total_amount"`
	}{
		OrderCode:   orderCode,
		TotalAmount: total,
	}

	return
}

func (p *PaymentSvcImpl) SimulationPayment(ctx context.Context, req SimulationPaymentReq) (resp models.DefaultResponse, err error) {
	{
		resp.Message = "Failed to simulate payment"
		resp.Code = http.StatusBadGateway
	}
	userData, ok := ctx.Value(middleware.UserData).(middleware.UserCtxReq)
	if !ok {
		slog.ErrorContext(ctx, "[PaymentSvc.SimulationPayment] error while get user data")
		resp.Message = "Failed to add product to cart"
		resp.Code = http.StatusUnauthorized
		return
	}

	payment, err := p.PaymentRepo.GetPaymentByOrderCode(ctx, int64(userData.UserID), req.OrderCode)
	if err != nil {
		slog.ErrorContext(ctx, "[PaymentSvc.SimulationPayment] error while GetPaymentByOrderCode err", "%v", err.Error())
		return
	}

	if payment.Status == "Settlement" {
		resp.Message = "Payment already settled"
		resp.Code = http.StatusBadRequest
		return
	}

	if req.Amount != payment.TotalAmount {
		resp.Message = "Invalid amount"
		resp.Code = http.StatusBadRequest
		return
	}

	err = p.PaymentRepo.UpdatePaymentStatus(ctx, int64(userData.UserID), req.OrderCode, "Settlement")
	if err != nil {
		slog.ErrorContext(ctx, "[PaymentSvc.SimulationPayment] error while UpdatePaymentStatus err", "%v", err.Error())
		return
	}

	resp.Message = "Payment simulated successfully"
	resp.Code = http.StatusOK

	return
}
