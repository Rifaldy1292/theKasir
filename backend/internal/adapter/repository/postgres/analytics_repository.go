package postgres

import (
	"context"

	"thekasir/internal/core/domain"
	"thekasir/internal/core/port"
)

type analyticsRepository struct {
	db *DB
}

func NewAnalyticsRepository(db *DB) port.AnalyticsRepository {
	return &analyticsRepository{db: db}
}

func (r *analyticsRepository) GetSummary(ctx context.Context, workspaceID string) (*domain.AnalyticsSummary, error) {
	// Coalesce to handle null when there are no transactions
	query := `
		SELECT 
			COALESCE(SUM(total_amount), 0) as total_revenue,
			COUNT(id) as total_transactions
		FROM transactions 
		WHERE workspace_id = $1 AND status = 'completed'
	`
	var summary domain.AnalyticsSummary
	err := r.db.Pool.QueryRow(ctx, query, workspaceID).Scan(
		&summary.TotalRevenue,
		&summary.TotalTransactions,
	)
	if err != nil {
		return nil, err
	}

	// Count active products
	qProd := `SELECT COUNT(id) FROM products WHERE workspace_id = $1`
	err = r.db.Pool.QueryRow(ctx, qProd, workspaceID).Scan(&summary.ActiveProducts)
	if err != nil {
		return nil, err
	}

	return &summary, nil
}

func (r *analyticsRepository) GetTopProducts(ctx context.Context, workspaceID string, limit int) ([]domain.TopProduct, error) {
	query := `
		SELECT 
			p.id, 
			p.name, 
			COALESCE(SUM(ti.quantity), 0) as total_sold,
			COALESCE(SUM(ti.subtotal), 0) as total_revenue
		FROM products p
		LEFT JOIN transaction_items ti ON p.id = ti.product_id
		LEFT JOIN transactions t ON ti.transaction_id = t.id AND t.status = 'completed'
		WHERE p.workspace_id = $1
		GROUP BY p.id, p.name
		ORDER BY total_sold DESC
		LIMIT $2
	`
	rows, err := r.db.Pool.Query(ctx, query, workspaceID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tops []domain.TopProduct
	for rows.Next() {
		var t domain.TopProduct
		if err := rows.Scan(&t.ProductID, &t.ProductName, &t.TotalSold, &t.TotalRevenue); err != nil {
			return nil, err
		}
		// Only include if actually sold something
		if t.TotalSold > 0 {
			tops = append(tops, t)
		}
	}
	return tops, nil
}

func (r *analyticsRepository) GetRecentTransactions(ctx context.Context, workspaceID string, limit int) ([]domain.Transaction, error) {
	query := `
		SELECT id, workspace_id, cashier_id, total_amount, payment_method, status, created_at
		FROM transactions
		WHERE workspace_id = $1
		ORDER BY created_at DESC
		LIMIT $2
	`
	rows, err := r.db.Pool.Query(ctx, query, workspaceID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var txs []domain.Transaction
	for rows.Next() {
		var t domain.Transaction
		if err := rows.Scan(&t.ID, &t.WorkspaceID, &t.CashierID, &t.TotalAmount, &t.PaymentMethod, &t.Status, &t.CreatedAt); err != nil {
			return nil, err
		}
		txs = append(txs, t)
	}
	return txs, nil
}
