package service

import (
	"be-shop/internal/app/models"
	"be-shop/internal/app/repo/postgres"
	"context"
	"log/slog"
	"math"
	"net/http"

	"go.uber.org/dig"
)

type (
	UpdatePriceReq struct {
		Price float64 `json:"price" validate:"required,gt=0,numeric"`
	}
	ProductSvc interface {
		CreateProduct(ctx context.Context, req models.Product) (resp models.DefaultResponse, err error)
		GetAllProduct(ctx context.Context, req models.PaginationRequest) (resp models.DefaultResponse, err error)
		GetProductByID(ctx context.Context, id int64) (resp models.DefaultResponse, err error)
		GetProductByCategoryID(ctx context.Context, id int64) (resp models.DefaultResponse, err error)
		UpdateProductPrice(ctx context.Context, id int64, req UpdatePriceReq) (resp models.DefaultResponse, err error)
		DeleteProduct(ctx context.Context, id int64) (resp models.DefaultResponse, err error)
		CreateCategory(ctx context.Context, name string) (resp models.DefaultResponse, err error)
	}

	ProductSvcImpl struct {
		dig.In

		ProductRepo  postgres.ProductRepo
		CategoryRepo postgres.CategoryRepo
	}
)

func NewProductSvc(impl ProductSvcImpl) ProductSvc {
	return &impl
}

func (p *ProductSvcImpl) CreateProduct(ctx context.Context, req models.Product) (resp models.DefaultResponse, err error) {
	{
		resp.Message = "Failed to create product"
		resp.Code = http.StatusBadGateway
	}

	id, err := p.ProductRepo.CreateProduct(ctx, req)
	if err != nil {
		slog.ErrorContext(ctx, "[ProductSvcImpl.CreateProduct] error while CreateProduct err", "%v", err.Error())
		return
	}

	resp.Message = "Product created successfully"
	resp.Code = http.StatusCreated
	resp.Data = struct {
		ID int `json:"id"`
	}{
		ID: id,
	}
	return
}

func (p *ProductSvcImpl) GetProductByID(ctx context.Context, id int64) (resp models.DefaultResponse, err error) {
	{
		resp.Message = "Failed to get product"
		resp.Code = http.StatusBadRequest
	}

	product, err := p.ProductRepo.GetProductByID(ctx, id)
	if err != nil {
		slog.ErrorContext(ctx, "[ProductSvcImpl.GetProductByID] error while GetProductByID err: %v", err.Error())
		return
	}

	resp.Message = "Product fetched successfully"
	resp.Code = http.StatusOK
	resp.Data = product
	return
}

func (p *ProductSvcImpl) GetAllProduct(ctx context.Context, req models.PaginationRequest) (resp models.DefaultResponse, err error) {
	{
		resp.Message = "Failed to get products"
		resp.Code = http.StatusBadGateway
		req.Page = (req.Page - 1) * req.Limit
	}
	totalItem, products, err := p.ProductRepo.GetAllProduct(ctx, req.Page, req.Limit)
	if err != nil {
		slog.ErrorContext(ctx, "[ProductSvcImpl.GetAllProduct] error while GetAllProduct err", "%v", err.Error())
		return
	}

	resp.Message = "Products fetched successfully"
	resp.Code = http.StatusOK
	resp.Data = models.DefaultPaginationResponseData{
		Results: products,
		DefaultMetaData: models.DefaultMetaData{
			Page:        uint(req.Page + 1),
			TotalPages:  uint(math.Ceil(float64(totalItem) / float64(req.Limit))),
			Limit:       uint(req.Limit),
			TotalItems:  uint(totalItem),
			HasNext:     req.Page < int(math.Ceil(float64(totalItem)/float64(req.Limit))),
			HasPrevious: req.Page > 1,
		},
	}
	return

}

func (p *ProductSvcImpl) UpdateProductPrice(ctx context.Context, id int64, req UpdatePriceReq) (resp models.DefaultResponse, err error) {
	{
		resp.Message = "Failed to update product price"
		resp.Code = http.StatusBadRequest
	}

	err = p.ProductRepo.UpdateProductPrice(ctx, id, req.Price)
	if err != nil {
		slog.ErrorContext(ctx, "[ProductSvcImpl.UpdateProductPrice] error while UpdateProductPrice err", "%v", err.Error())
		return
	}

	resp.Message = "Product price updated successfully"
	resp.Code = http.StatusOK
	return
}

func (p *ProductSvcImpl) DeleteProduct(ctx context.Context, id int64) (resp models.DefaultResponse, err error) {
	{
		resp.Message = "Failed to delete product"
		resp.Code = http.StatusBadRequest
	}

	err = p.ProductRepo.SoftDeleteProduct(ctx, id)
	if err != nil {
		slog.ErrorContext(ctx, "[ProductSvcImpl.DeleteProduct] error while SoftDeleteProduct err", "%v", err.Error())
		return
	}

	resp.Message = "Product deleted successfully"
	resp.Code = http.StatusOK
	return
}

func (p *ProductSvcImpl) GetProductByCategoryID(ctx context.Context, id int64) (resp models.DefaultResponse, err error) {
	{
		resp.Message = "Failed to get product"
		resp.Code = http.StatusBadRequest
	}

	product, err := p.ProductRepo.GetProductByCategoryID(ctx, id)
	if err != nil {
		slog.ErrorContext(ctx, "[ProductSvcImpl.GetProductByCategoryID] error while GetProductByCategoryID err", "%v", err.Error())
		return
	}

	resp.Message = "Product fetched successfully"
	resp.Code = http.StatusOK
	resp.Data = product
	return
}

func (p *ProductSvcImpl) CreateCategory(ctx context.Context, name string) (resp models.DefaultResponse, err error) {
	{
		resp.Message = "Failed to create category"
		resp.Code = http.StatusBadRequest
	}

	id, err := p.CategoryRepo.CreateCategory(ctx, name)
	if err != nil {
		slog.ErrorContext(ctx, "[ProductSvcImpl.CreateCategory] error while CreateCategory err", "%v", err.Error())
		return
	}

	resp.Message = "Category created successfully"
	resp.Code = http.StatusCreated
	resp.Data = struct {
		ID int `json:"id"`
	}{
		ID: id,
	}
	return
}
