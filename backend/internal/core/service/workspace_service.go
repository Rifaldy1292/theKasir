package service

import (
	"context"
	"thekasir/internal/core/domain"
	"thekasir/internal/core/port"
	"time"
)

type workspaceService struct {
	workspaceRepo  port.WorkspaceRepository
	membershipRepo port.MembershipRepository
}

func NewWorkspaceService(wsRepo port.WorkspaceRepository, memRepo port.MembershipRepository) port.WorkspaceService {
	return &workspaceService{
		workspaceRepo:  wsRepo,
		membershipRepo: memRepo,
	}
}

func (s *workspaceService) CreateWorkspace(ctx context.Context, name, businessType, userID string) (*domain.Workspace, error) {
	// Simple ID generation
	wsID := "ws_" + time.Now().Format("20060102150405")

	ws := &domain.Workspace{
		ID:           wsID,
		Name:         name,
		BusinessType: businessType,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := s.workspaceRepo.CreateWorkspace(ctx, ws); err != nil {
		return nil, err
	}

	// Create owner membership
	memID := "mem_" + time.Now().Format("20060102150405")
	membership := &domain.Membership{
		ID:          memID,
		UserID:      userID,
		WorkspaceID: wsID,
		Role:        domain.RoleOwner,
		JoinedAt:    time.Now(),
	}

	if err := s.membershipRepo.CreateMembership(ctx, membership); err != nil {
		return nil, err
	}

	return ws, nil
}

func (s *workspaceService) GetUserWorkspaces(ctx context.Context, userID string) ([]domain.Workspace, error) {
	// In a real app, you would query workspaces through memberships.
	// For simplicity in Phase 1, we return empty list if not fully implemented in repo.
	// This requires joining memberships and workspaces.
	return nil, nil // TODO: implement join in repo
}
