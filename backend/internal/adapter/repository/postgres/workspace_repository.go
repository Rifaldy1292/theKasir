package postgres

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"thekasir/internal/core/domain"
	"thekasir/internal/core/port"
)

type workspaceRepository struct {
	db *DB
}

func NewWorkspaceRepository(db *DB) port.WorkspaceRepository {
	return &workspaceRepository{db: db}
}

func (r *workspaceRepository) CreateWorkspace(ctx context.Context, workspace *domain.Workspace) error {
	query := `
		INSERT INTO workspaces (id, name, business_type, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.db.Pool.Exec(ctx, query,
		workspace.ID,
		workspace.Name,
		workspace.BusinessType,
		workspace.CreatedAt,
		workspace.UpdatedAt,
	)
	return err
}

func (r *workspaceRepository) GetWorkspaceByID(ctx context.Context, id string) (*domain.Workspace, error) {
	query := `
		SELECT id, name, business_type, created_at, updated_at
		FROM workspaces
		WHERE id = $1
	`
	var w domain.Workspace
	err := r.db.Pool.QueryRow(ctx, query, id).Scan(
		&w.ID,
		&w.Name,
		&w.BusinessType,
		&w.CreatedAt,
		&w.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &w, nil
}
