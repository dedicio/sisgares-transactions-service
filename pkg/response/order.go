package response

import (
	"net/http"

	"github.com/dedicio/sisgares-transactions-service/internal/dto"
)

type OrderResponse struct {
	*dto.OrderResponseDto
}

func (pr *OrderResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewOrderResponse(product *dto.OrderResponseDto) *OrderResponse {
	return &OrderResponse{product}
}

type OrdersResponse struct {
	Orders []*dto.OrderResponseDto
}

func (pr *OrdersResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewOrdersResponse(products []*dto.OrderResponseDto) *OrdersResponse {
	return &OrdersResponse{products}
}
