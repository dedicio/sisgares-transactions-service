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
	Items []*dto.OrderResponseDto
}

func (pr *OrdersResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewOrdersResponse(products []*dto.OrderResponseDto) *OrdersResponse {
	return &OrdersResponse{products}
}

type OrderItemResponse struct {
	*dto.OrderItemResponseDto
}

func (pr *OrderItemResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewOrderItemResponse(product *dto.OrderItemResponseDto) *OrderItemResponse {
	return &OrderItemResponse{product}
}
