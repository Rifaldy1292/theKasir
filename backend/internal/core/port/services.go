package port

import (
	"context"
	"thekasir/internal/core/domain"
)

type AuthService interface {
	Register(ctx context.Context, email, name, password string) (*domain.User, error)
	Login(ctx context.Context, email, password string) (string, *domain.User, error)
}

type WorkspaceService interface {
	CreateWorkspace(ctx context.Context, name, businessType, userID string) (*domain.Workspace, error)
	GetUserWorkspaces(ctx context.Context, userID string) ([]domain.Workspace, error)
}
