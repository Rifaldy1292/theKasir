package port

import (
	"context"
	"thekasir/internal/core/domain"
)

type PosRepository interface {
	// Product operations
	CreateProduct(ctx context.Context, product *domain.Product) error
	GetProductsByWorkspace(ctx context.Context, workspaceID string) ([]domain.Product, error)
	GetProductByID(ctx context.Context, id, workspaceID string) (*domain.Product, error)
	UpdateProductStock(ctx context.Context, productID string, quantity int) error

	// Transaction operations
	CreateTransaction(ctx context.Context, tx *domain.Transaction, items []domain.TransactionItem) error
}

type PosService interface {
	CreateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error)
	GetProducts(ctx context.Context, workspaceID string) ([]domain.Product, error)
	Checkout(ctx context.Context, req domain.CheckoutRequest) (*domain.Transaction, error)
}
