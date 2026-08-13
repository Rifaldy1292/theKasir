package port

import (
	"context"
	"thekasir/internal/core/domain"
)

type AnalyticsRepository interface {
	GetSummary(ctx context.Context, workspaceID string) (*domain.AnalyticsSummary, error)
	GetTopProducts(ctx context.Context, workspaceID string, limit int) ([]domain.TopProduct, error)
	GetRecentTransactions(ctx context.Context, workspaceID string, limit int) ([]domain.Transaction, error)
}

type AnalyticsService interface {
	GetDashboardData(ctx context.Context, workspaceID string) (*domain.DashboardResponse, error)
}
