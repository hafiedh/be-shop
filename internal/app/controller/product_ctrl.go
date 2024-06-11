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
	ProductCtrl interface {
		CreateProduct(ec echo.Context) error
		GetProductByID(ec echo.Context) error
		GetProductsByCategoryID(ec echo.Context) error
		GetAllProduct(ec echo.Context) error
		UpdateProductPrice(ec echo.Context) error
		DeleteProduct(ec echo.Context) error

		CreateCategory(ec echo.Context) error
	}

	ProductCtrlImpl struct {
		dig.In

		ProductSvc service.ProductSvc
	}
)

func NewProductCtrl(impl ProductCtrlImpl) ProductCtrl {
	return &impl
}

func (m *ProductCtrlImpl) CreateProduct(ec echo.Context) error {
	Recover()
	ctx := ec.Request().Context()
	var req models.Product

	if err := ec.Bind(&req); err != nil {
		slog.Error("CreateProduct - Invalid request body", err)
		return ec.JSON(http.StatusBadRequest, models.DefaultResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	validate := utils.Validate

	err := validate.Struct(req)
	if err != nil {
		slog.Error("CreateModule - validation error", err)
		errors := err.(validator.ValidationErrors)
		return ec.JSON(http.StatusBadRequest, models.DefaultResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
			Error:   errors.Error(),
		})
	}

	resp, err := m.ProductSvc.CreateProduct(ctx, req)
	if err != nil {
		slog.Error("CreateProduct - error while creating product", err)
		return ec.JSON(http.StatusInternalServerError, models.DefaultResponse{
			Code:    resp.Code,
			Message: "Failed to create product",
			Error:   err.(validator.ValidationErrors).Error(),
		})
	}

	return ec.JSON(resp.Code, resp)
}

func (m *ProductCtrlImpl) GetProductByID(ec echo.Context) error {
	Recover()
	ctx := ec.Request().Context()
	id := ec.Param("id")

	idConv, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		slog.Error("GetProductByID - error while converting id", err)
		return ec.JSON(http.StatusBadRequest, models.DefaultResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	resp, err := m.ProductSvc.GetProductByID(ctx, idConv)
	if err != nil {
		slog.Error("GetProductByID - error while getting product", err)
		return ec.JSON(http.StatusInternalServerError, models.DefaultResponse{
			Code:    resp.Code,
			Message: "Failed to get product",
			Error:   err.(validator.ValidationErrors).Error(),
		})
	}

	return ec.JSON(resp.Code, resp)
}

func (m *ProductCtrlImpl) GetAllProduct(ec echo.Context) error {
	Recover()
	ctx := ec.Request().Context()

	var req models.PaginationRequest

	if err := ec.Bind(&req); err != nil {
		slog.Error("GetAllProduct - Invalid request body", err)
		return ec.JSON(http.StatusBadRequest, models.DefaultResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	switch {
	case req.Page == 0 && req.Limit == 0:
		req.SetDefaults()
	case req.Page == 0:
		req.SetDefaultPage()
	case req.Limit == 0:
		req.SetDefaultLimit()
	}

	resp, err := m.ProductSvc.GetAllProduct(ctx, req)
	if err != nil {
		slog.Error("GetAllProduct - error while getting all products", err)
		return ec.JSON(http.StatusInternalServerError, models.DefaultResponse{
			Code:    resp.Code,
			Message: "Failed to get products",
			Error:   err.(validator.ValidationErrors).Error(),
		})
	}

	return ec.JSON(resp.Code, resp)
}

func (m *ProductCtrlImpl) UpdateProductPrice(ec echo.Context) error {
	Recover()
	ctx := ec.Request().Context()

	id := ec.Param("id")

	idConv, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		slog.Error("UpdateProductPrice - error while converting id", err)
		return ec.JSON(http.StatusBadRequest, models.DefaultResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	var req service.UpdatePriceReq

	if err := ec.Bind(&req); err != nil {
		slog.Error("UpdateProductPrice - Invalid request body", err)
		return ec.JSON(http.StatusBadRequest, models.DefaultResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	validate := utils.Validate

	err = validate.Struct(req)
	if err != nil {
		slog.Error("UpdateProductPrice - validation error", err)
		errors := err.(validator.ValidationErrors)
		return ec.JSON(http.StatusBadRequest, models.DefaultResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
			Error:   errors.Error(),
		})
	}

	resp, err := m.ProductSvc.UpdateProductPrice(ctx, idConv, req)
	if err != nil {
		slog.Error("UpdateProductPrice - error while updating product price", err)
		return ec.JSON(http.StatusInternalServerError, models.DefaultResponse{
			Code:    resp.Code,
			Message: "Failed to update product price",
			Error:   err.(validator.ValidationErrors).Error(),
		})
	}

	return ec.JSON(resp.Code, resp)

}

func (m *ProductCtrlImpl) DeleteProduct(ec echo.Context) error {
	Recover()
	ctx := ec.Request().Context()

	id := ec.Param("id")

	idConv, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		slog.Error("DeleteProduct - error while converting id", err)
		return ec.JSON(http.StatusBadRequest, models.DefaultResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	resp, err := m.ProductSvc.DeleteProduct(ctx, idConv)
	if err != nil {
		slog.Error("DeleteProduct - error while deleting product", err)
		return ec.JSON(http.StatusInternalServerError, models.DefaultResponse{
			Code:    resp.Code,
			Message: "Failed to delete product",
			Error:   err.(validator.ValidationErrors).Error(),
		})
	}

	return ec.JSON(resp.Code, resp)
}

func (m *ProductCtrlImpl) GetProductsByCategoryID(ec echo.Context) error {
	Recover()
	ctx := ec.Request().Context()

	id := ec.Param("id")

	idConv, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		slog.Error("GetProductsByCategoryID - error while converting id", err)
		return ec.JSON(http.StatusBadRequest, models.DefaultResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	resp, err := m.ProductSvc.GetProductByCategoryID(ctx, idConv)
	if err != nil {
		slog.Error("GetProductsByCategoryID - error while getting product by category id", err)
		return ec.JSON(http.StatusInternalServerError, models.DefaultResponse{
			Code:    resp.Code,
			Message: "Failed to get product by category id",
			Error:   err.(validator.ValidationErrors).Error(),
		})
	}

	return ec.JSON(resp.Code, resp)
}

func (m *ProductCtrlImpl) CreateCategory(ec echo.Context) error {
	Recover()
	ctx := ec.Request().Context()

	var req models.Category

	if err := ec.Bind(&req); err != nil {
		slog.Error("CreateCategory - Invalid request body", err)
		return ec.JSON(http.StatusBadRequest, models.DefaultResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	validate := utils.Validate

	err := validate.Struct(req)
	if err != nil {
		slog.Error("CreateCategory - validation error", err)
		errors := err.(validator.ValidationErrors)
		return ec.JSON(http.StatusBadRequest, models.DefaultResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
			Error:   errors.Error(),
		})
	}

	resp, err := m.ProductSvc.CreateCategory(ctx, req.Name)
	if err != nil {
		slog.Error("CreateCategory - error while creating category", err)
		return ec.JSON(http.StatusInternalServerError, models.DefaultResponse{
			Code:    resp.Code,
			Message: "Failed to create category",
			Error:   err.(validator.ValidationErrors).Error(),
		})
	}

	return ec.JSON(resp.Code, resp)
}
