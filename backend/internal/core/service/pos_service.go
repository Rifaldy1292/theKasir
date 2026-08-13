package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"thekasir/internal/core/domain"
	"thekasir/internal/core/port"
)

type posService struct {
	posRepo port.PosRepository
}

func NewPosService(repo port.PosRepository) port.PosService {
	return &posService{posRepo: repo}
}

func (s *posService) CreateProduct(ctx context.Context, p *domain.Product) (*domain.Product, error) {
	if p.Name == "" || p.Price < 0 {
		return nil, errors.New("invalid product data")
	}

	p.ID = "prd_" + time.Now().Format("20060102150405")
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	err := s.posRepo.CreateProduct(ctx, p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (s *posService) GetProducts(ctx context.Context, workspaceID string) ([]domain.Product, error) {
	return s.posRepo.GetProductsByWorkspace(ctx, workspaceID)
}

func (s *posService) Checkout(ctx context.Context, req domain.CheckoutRequest) (*domain.Transaction, error) {
	if len(req.Items) == 0 {
		return nil, errors.New("cart is empty")
	}

	tx := &domain.Transaction{
		ID:            "tx_" + time.Now().Format("20060102150405"),
		WorkspaceID:   req.WorkspaceID,
		CashierID:     req.CashierID,
		PaymentMethod: req.PaymentMethod,
		Status:        "completed",
		CreatedAt:     time.Now(),
	}

	var totalAmount float64
	var txItems []domain.TransactionItem

	// Validate items and calculate total
	for i, item := range req.Items {
		product, err := s.posRepo.GetProductByID(ctx, item.ProductID, req.WorkspaceID)
		if err != nil {
			return nil, fmt.Errorf("failed to get product %s: %w", item.ProductID, err)
		}
		if product == nil {
			return nil, fmt.Errorf("product %s not found in workspace", item.ProductID)
		}
		if product.Stock < item.Quantity {
			return nil, fmt.Errorf("insufficient stock for product %s", product.Name)
		}

		subtotal := product.Price * float64(item.Quantity)
		totalAmount += subtotal

		txItem := domain.TransactionItem{
			ID:            fmt.Sprintf("txi_%s_%d", tx.ID, i),
			TransactionID: tx.ID,
			ProductID:     product.ID,
			Quantity:      item.Quantity,
			UnitPrice:     product.Price,
			Subtotal:      subtotal,
			CreatedAt:     time.Now(),
		}
		txItems = append(txItems, txItem)
	}

	tx.TotalAmount = totalAmount

	// Save transaction (includes stock reduction)
	err := s.posRepo.CreateTransaction(ctx, tx, txItems)
	if err != nil {
		return nil, err
	}

	return tx, nil
}
