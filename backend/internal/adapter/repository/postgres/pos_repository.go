package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"thekasir/internal/core/domain"
	"thekasir/internal/core/port"
)

type posRepository struct {
	db *DB
}

func NewPosRepository(db *DB) port.PosRepository {
	return &posRepository{db: db}
}

func (r *posRepository) CreateProduct(ctx context.Context, p *domain.Product) error {
	query := `
		INSERT INTO products (id, workspace_id, category_id, name, price, cost, stock, image_url, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`
	_, err := r.db.Pool.Exec(ctx, query,
		p.ID, p.WorkspaceID, p.CategoryID, p.Name, p.Price, p.Cost, p.Stock, p.ImageURL, p.CreatedAt, p.UpdatedAt,
	)
	return err
}

func (r *posRepository) GetProductsByWorkspace(ctx context.Context, workspaceID string) ([]domain.Product, error) {
	query := `
		SELECT id, workspace_id, category_id, name, price, cost, stock, image_url, created_at, updated_at
		FROM products
		WHERE workspace_id = $1
		ORDER BY name ASC
	`
	rows, err := r.db.Pool.Query(ctx, query, workspaceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		var p domain.Product
		err := rows.Scan(
			&p.ID, &p.WorkspaceID, &p.CategoryID, &p.Name, &p.Price, &p.Cost, &p.Stock, &p.ImageURL, &p.CreatedAt, &p.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (r *posRepository) GetProductByID(ctx context.Context, id, workspaceID string) (*domain.Product, error) {
	query := `
		SELECT id, workspace_id, category_id, name, price, cost, stock, image_url, created_at, updated_at
		FROM products
		WHERE id = $1 AND workspace_id = $2
	`
	var p domain.Product
	err := r.db.Pool.QueryRow(ctx, query, id, workspaceID).Scan(
		&p.ID, &p.WorkspaceID, &p.CategoryID, &p.Name, &p.Price, &p.Cost, &p.Stock, &p.ImageURL, &p.CreatedAt, &p.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

func (r *posRepository) UpdateProductStock(ctx context.Context, productID string, quantity int) error {
	query := `UPDATE products SET stock = stock - $1, updated_at = NOW() WHERE id = $2`
	_, err := r.db.Pool.Exec(ctx, query, quantity, productID)
	return err
}

func (r *posRepository) CreateTransaction(ctx context.Context, tx *domain.Transaction, items []domain.TransactionItem) error {
	// Begin DB Transaction
	dbtx, err := r.db.Pool.Begin(ctx)
	if err != nil {
		return err
	}
	// Defer a rollback in case anything fails. If commit is successful, rollback is a no-op.
	defer dbtx.Rollback(ctx)

	// Insert Transaction Header
	qTx := `
		INSERT INTO transactions (id, workspace_id, cashier_id, total_amount, payment_method, status, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err = dbtx.Exec(ctx, qTx,
		tx.ID, tx.WorkspaceID, tx.CashierID, tx.TotalAmount, tx.PaymentMethod, tx.Status, tx.CreatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to insert transaction: %w", err)
	}

	// Insert Transaction Items and Update Stock
	qItem := `
		INSERT INTO transaction_items (id, transaction_id, product_id, quantity, unit_price, subtotal, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	qStock := `UPDATE products SET stock = stock - $1 WHERE id = $2`

	for _, item := range items {
		// Insert item
		_, err = dbtx.Exec(ctx, qItem,
			item.ID, item.TransactionID, item.ProductID, item.Quantity, item.UnitPrice, item.Subtotal, item.CreatedAt,
		)
		if err != nil {
			return fmt.Errorf("failed to insert transaction item: %w", err)
		}

		// Update stock
		_, err = dbtx.Exec(ctx, qStock, item.Quantity, item.ProductID)
		if err != nil {
			return fmt.Errorf("failed to update stock: %w", err)
		}
	}

	// Commit Transaction
	return dbtx.Commit(ctx)
}
