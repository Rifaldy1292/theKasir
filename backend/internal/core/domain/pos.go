package domain

import (
	"time"
)

type Category struct {
	ID          string    `json:"id"`
	WorkspaceID string    `json:"workspace_id"`
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Product struct {
	ID          string    `json:"id"`
	WorkspaceID string    `json:"workspace_id"`
	CategoryID  *string   `json:"category_id"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Cost        float64   `json:"cost"`
	Stock       int       `json:"stock"`
	ImageURL    string    `json:"image_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Transaction struct {
	ID            string    `json:"id"`
	WorkspaceID   string    `json:"workspace_id"`
	CashierID     string    `json:"cashier_id"`
	TotalAmount   float64   `json:"total_amount"`
	PaymentMethod string    `json:"payment_method"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
}

type TransactionItem struct {
	ID            string    `json:"id"`
	TransactionID string    `json:"transaction_id"`
	ProductID     string    `json:"product_id"`
	Quantity      int       `json:"quantity"`
	UnitPrice     float64   `json:"unit_price"`
	Subtotal      float64   `json:"subtotal"`
	CreatedAt     time.Time `json:"created_at"`
}

type CheckoutRequest struct {
	WorkspaceID   string `json:"-"` // injected from route
	CashierID     string `json:"-"` // injected from auth middleware
	PaymentMethod string `json:"payment_method"`
	Items         []struct {
		ProductID string `json:"product_id"`
		Quantity  int    `json:"quantity"`
	} `json:"items"`
}
