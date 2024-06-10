package postgres

import (
	"be-shop/internal/app/models"
	"be-shop/internal/app/repo/postgres/queries"
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	"go.uber.org/dig"
)

type (
	ProductRepo interface {
		CreateProduct(ctx context.Context, req models.Product) (id int, err error)
		GetProductByID(ctx context.Context, id int64) (product models.Product, err error)
		GetAllProduct(ctx context.Context, page, limit int) (totalItem int, products []models.Product, err error)
		GetProductByCategoryID(ctx context.Context, id int64) (resp []models.Product, err error)
		UpdateProductPrice(ctx context.Context, id int64, price float64) (err error)
		SoftDeleteProduct(ctx context.Context, id int64) (err error)
	}

	ProductRepoImpl struct {
		dig.In

		*sql.DB
	}
)

func NewProductRepo(impl ProductRepoImpl) ProductRepo {
	return &impl
}

func (p *ProductRepoImpl) CreateProduct(ctx context.Context, req models.Product) (id int, err error) {
	_, err = p.QueryContext(ctx, queries.QueryCreateProduct, req.Name, req.CategoryID, req.Price)
	if err != nil {
		slog.ErrorContext(ctx, fmt.Sprintf("[ProductRepoImpl.CreateProduct] error while CreateProduct err: %v", err.Error()))
		return id, err
	}

	return id, nil
}

func (p *ProductRepoImpl) GetProductByID(ctx context.Context, id int64) (product models.Product, err error) {
	row := p.QueryRowContext(ctx, queries.QueryGetProductByID, id)
	err = row.Scan(&product.ID, &product.Name, &product.CategoryID, &product.Price, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		slog.ErrorContext(ctx, fmt.Sprintf("error while GetProductByID err: %v", err.Error()))
		return
	}
	return
}

func (p *ProductRepoImpl) GetAllProduct(ctx context.Context, page, limit int) (totalItem int, products []models.Product, err error) {
	rows, err := p.QueryContext(ctx, queries.QueryGetAllProducts, limit, page)
	if err != nil {
		slog.ErrorContext(ctx, fmt.Sprintf("[ProductRepoImpl.GetAllProduct] error while GetAllProduct err: %v", err.Error()))
		return
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		err = rows.Scan(&totalItem, &product.ID, &product.Name, &product.CategoryID, &product.Price, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			slog.ErrorContext(ctx, fmt.Sprintf("[ProductRepoImpl.GetAllProduct] error while GetAllProduct err: %v", err.Error()))
			return
		}
		products = append(products, product)
	}

	if products == nil {
		products = make([]models.Product, 0)
		slog.InfoContext(ctx, fmt.Sprintf("[ProductRepoImpl.GetAllProduct] products: %v", products))
		return
	}

	slog.InfoContext(ctx, fmt.Sprintf("[ProductRepoImpl.GetAllProduct] totalItem: %v", totalItem))

	return
}

func (p *ProductRepoImpl) UpdateProductPrice(ctx context.Context, id int64, price float64) (err error) {
	_, err = p.ExecContext(ctx, queries.QueryUpdateProductPrice, price, id)
	if err != nil {
		slog.ErrorContext(ctx, fmt.Sprintf("[ProductRepoImpl.UpdateProductPrice] error while UpdateProductPrice err: %v", err.Error()))
		return err
	}
	return
}

func (p *ProductRepoImpl) SoftDeleteProduct(ctx context.Context, id int64) (err error) {
	_, err = p.ExecContext(ctx, queries.QueryDeleteProduct, id)
	if err != nil {
		slog.ErrorContext(ctx, fmt.Sprintf("[ProductRepoImpl.SoftDeleteProduct] error while SoftDeleteProduct err: %v", err.Error()))
		return err
	}
	return
}

func (p *ProductRepoImpl) GetProductByCategoryID(ctx context.Context, id int64) (products []models.Product, err error) {
	rows, err := p.QueryContext(ctx, queries.QueryGetProductByCategoryID, id)
	if err != nil {
		slog.ErrorContext(ctx, fmt.Sprintf("[ProductRepoImpl.GetProductByCategoryID] error while GetProductByCategoryID err: %v", err.Error()))
		return
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		err = rows.Scan(&product.ID, &product.Name, &product.CategoryID, &product.Price, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			slog.ErrorContext(ctx, fmt.Sprintf("[ProductRepoImpl.GetProductByCategoryID] error while GetProductByCategoryID err: %v", err.Error()))
			return
		}
		products = append(products, product)
	}

	if products == nil {
		products = make([]models.Product, 0)
		slog.InfoContext(ctx, fmt.Sprintf("[ProductRepoImpl.GetProductByCategoryID] products: %v", products))
		return
	}

	return
}
