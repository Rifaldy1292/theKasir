package domain

import (
	"time"
)

type Role string

const (
	RoleOwner Role = "owner"
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type Membership struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	WorkspaceID string    `json:"workspace_id"`
	Role        Role      `json:"role"`
	JoinedAt    time.Time `json:"joined_at"`
}
