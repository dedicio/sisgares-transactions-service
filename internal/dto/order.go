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
	OrderID   string  `json:"order_id"`
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
}

type OrderResponseDto struct {
	ID            string         `json:"id"`
	Items         []OrderItemDto `json:"items"`
	Discount      float64        `json:"discount"`
	Status        string         `json:"status"`
	PaymentMethod string         `json:"payment_method"`
	TotalPrice    float64        `json:"total_price"`
}

type OrderItemOutputDto struct {
	ID        string  `json:"id"`
	OrderID   string  `json:"order_id"`
	ProductID string  `json:"product_id"`
	Quantity  int64   `json:"quantity"`
	Price     float64 `json:"price"`
	CompanyId string  `json:"company_id"`
}

type OrderItemResponseDto struct {
	ID        string  `json:"id"`
	OrderID   string  `json:"order_id"`
	ProductID string  `json:"product_id"`
	Quantity  int64   `json:"quantity"`
	Price     float64 `json:"price"`
}
