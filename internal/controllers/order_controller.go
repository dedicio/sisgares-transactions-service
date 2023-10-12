package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dedicio/sisgares-transactions-service/internal/dto"
	"github.com/dedicio/sisgares-transactions-service/internal/entity"
	usecase "github.com/dedicio/sisgares-transactions-service/internal/usecase"
	httpResponsePkg "github.com/dedicio/sisgares-transactions-service/pkg/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type OrderController struct {
	Repository entity.OrderRepository
	Publisher  entity.OrderPublisher
}

func NewOrderController(
	orderRepository entity.OrderRepository,
	orderPublisher entity.OrderPublisher,
) *OrderController {
	return &OrderController{
		Repository: orderRepository,
		Publisher:  orderPublisher,
	}
}

func (lc *OrderController) FindAll(w http.ResponseWriter, r *http.Request) {
	companyID := r.Header.Get("X-Company-ID")
	orders, err := usecase.NewListOrdersUseCase(lc.Repository).Execute(companyID)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	render.Render(w, r, httpResponsePkg.NewOrdersResponse(orders))
}

func (lc *OrderController) FindById(w http.ResponseWriter, r *http.Request) {
	orderId := chi.URLParam(r, "id")
	order, err := usecase.NewFindOrderByIdUseCase(lc.Repository).Execute(orderId)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrNotFound(err, "Pedido"))
		return
	}

	render.Render(w, r, httpResponsePkg.NewOrderResponse(order))
}

func (lc *OrderController) Create(w http.ResponseWriter, r *http.Request) {
	companyID := r.Header.Get("X-Company-ID")
	payload := json.NewDecoder(r.Body)
	order := dto.OrderDto{}
	err := payload.Decode(&order)
	fmt.Println("order", order)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInvalidRequest(err))
		return
	}

	order.CompanyId = companyID
	orderSaved, err := usecase.NewCreateOrderUseCase(lc.Repository).Execute(order)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	output := &dto.OrderResponseDto{
		ID:         orderSaved.ID,
		Discount:   orderSaved.Discount,
		Items:      orderSaved.Items,
		Status:     orderSaved.Status,
		TotalPrice: orderSaved.TotalPrice,
		CreatedAt:  orderSaved.CreatedAt,
		UpdatedAt:  orderSaved.UpdatedAt,
	}

	render.Render(w, r, httpResponsePkg.NewOrderResponse(output))
}

func (lc *OrderController) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	orderId := chi.URLParam(r, "id")
	orderStatus := chi.URLParam(r, "status")
	err := usecase.NewUpdateOrderStatusUseCase(lc.Repository).Execute(orderId, orderStatus)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	output := &dto.OrderResponseDto{
		ID:     orderId,
		Status: orderStatus,
	}

	render.Render(w, r, httpResponsePkg.NewOrderResponse(output))
}
