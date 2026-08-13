package port

import (
	"context"
	"thekasir/internal/core/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	GetUserByID(ctx context.Context, id string) (*domain.User, error)
}

type WorkspaceRepository interface {
	CreateWorkspace(ctx context.Context, workspace *domain.Workspace) error
	GetWorkspaceByID(ctx context.Context, id string) (*domain.Workspace, error)
}

type MembershipRepository interface {
	CreateMembership(ctx context.Context, membership *domain.Membership) error
	GetUserMemberships(ctx context.Context, userID string) ([]domain.Membership, error)
	GetMembership(ctx context.Context, userID, workspaceID string) (*domain.Membership, error)
}
