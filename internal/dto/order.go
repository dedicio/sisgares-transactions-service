package dto

type OrderDto struct {
	ID            string         `json:"id"`
	Items         []OrderItemDto `json:"items"`
	Discount      float64        `json:"discount"`
	Status        string         `json:"status"`
	PaymentMethod string         `json:"payment_method"`
	TotalPrice    float64        `json:"total_price"`
	CompanyId     string         `json:"company_id"`
}

type OrderItemDto struct {
	ID        string  `json:"id"`
	ProductID string  `json:"product_id"`
	Quantity  int64   `json:"quantity"`
	Price     float64 `json:"price"`
}

type OrderOutputDto struct {
	ID            string         `json:"id"`
	Items         []OrderItemDto `json:"items"`
	Discount      float64        `json:"discount"`
	Status        string         `json:"status"`
	PaymentMethod string         `json:"payment_method"`
	TotalPrice    float64        `json:"total_price"`
	CompanyId     string         `json:"company_id"`
	CreatedAt     string         `json:"created_at"`
	UpdatedAt     string         `json:"updated_at"`
}

type OrderResponseDto struct {
	ID            string         `json:"id"`
	Items         []OrderItemDto `json:"items"`
	Discount      float64        `json:"discount"`
	Status        string         `json:"status"`
	PaymentMethod string         `json:"payment_method"`
	TotalPrice    float64        `json:"total_price"`
	CreatedAt     string         `json:"created_at"`
	UpdatedAt     string         `json:"updated_at"`
}
