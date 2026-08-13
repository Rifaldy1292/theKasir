package domain

type AnalyticsSummary struct {
	TotalRevenue      float64 `json:"total_revenue"`
	TotalTransactions int     `json:"total_transactions"`
	ActiveProducts    int     `json:"active_products"`
}

type TopProduct struct {
	ProductID   string  `json:"product_id"`
	ProductName string  `json:"product_name"`
	TotalSold   int     `json:"total_sold"`
	TotalRevenue float64 `json:"total_revenue"`
}

type DashboardResponse struct {
	Summary     AnalyticsSummary `json:"summary"`
	TopProducts []TopProduct     `json:"top_products"`
	RecentTx    []Transaction    `json:"recent_transactions"`
}
