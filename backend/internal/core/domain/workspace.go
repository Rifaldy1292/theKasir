package domain

import (
	"time"
)

type Workspace struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	BusinessType string    `json:"business_type"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
