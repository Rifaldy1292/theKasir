package service

import (
	"context"
	"thekasir/internal/core/domain"
	"thekasir/internal/core/port"
)

type analyticsService struct {
	repo port.AnalyticsRepository
}

func NewAnalyticsService(repo port.AnalyticsRepository) port.AnalyticsService {
	return &analyticsService{repo: repo}
}

func (s *analyticsService) GetDashboardData(ctx context.Context, workspaceID string) (*domain.DashboardResponse, error) {
	summary, err := s.repo.GetSummary(ctx, workspaceID)
	if err != nil {
		return nil, err
	}

	topProducts, err := s.repo.GetTopProducts(ctx, workspaceID, 5) // top 5
	if err != nil {
		return nil, err
	}

	recentTx, err := s.repo.GetRecentTransactions(ctx, workspaceID, 5) // 5 most recent
	if err != nil {
		return nil, err
	}

	// Prevent null slices in JSON
	if topProducts == nil {
		topProducts = []domain.TopProduct{}
	}
	if recentTx == nil {
		recentTx = []domain.Transaction{}
	}

	return &domain.DashboardResponse{
		Summary:     *summary,
		TopProducts: topProducts,
		RecentTx:    recentTx,
	}, nil
}
