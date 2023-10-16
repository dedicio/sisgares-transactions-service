package controllers

import (
	"encoding/json"
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

func (oc *OrderController) FindAll(w http.ResponseWriter, r *http.Request) {
	companyID := r.Header.Get("X-Company-ID")
	orders, err := usecase.NewListOrdersUseCase(oc.Repository).Execute(companyID)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	render.Render(w, r, httpResponsePkg.NewOrdersResponse(orders))
}

func (oc *OrderController) FindById(w http.ResponseWriter, r *http.Request) {
	orderId := chi.URLParam(r, "id")
	order, err := usecase.NewFindOrderByIdUseCase(oc.Repository).Execute(orderId)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrNotFound(err, "Pedido"))
		return
	}

	render.Render(w, r, httpResponsePkg.NewOrderResponse(order))
}

func (oc *OrderController) Create(w http.ResponseWriter, r *http.Request) {
	companyID := r.Header.Get("X-Company-ID")
	payload := json.NewDecoder(r.Body)
	order := dto.OrderDto{}
	err := payload.Decode(&order)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInvalidRequest(err))
		return
	}

	order.CompanyId = companyID
	orderSaved, err := usecase.NewCreateOrderUseCase(oc.Repository).Execute(order)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	output := &dto.OrderResponseDto{
		ID:            orderSaved.ID,
		Discount:      orderSaved.Discount,
		Items:         orderSaved.Items,
		Status:        orderSaved.Status,
		PaymentMethod: orderSaved.PaymentMethod,
		TotalPrice:    orderSaved.TotalPrice,
	}

	go usecase.NewPublishCreateOrderUseCase(oc.Publisher).Execute(orderSaved)

	render.Render(w, r, httpResponsePkg.NewOrderResponse(output))
}

func (oc *OrderController) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	orderId := chi.URLParam(r, "id")
	orderStatus := chi.URLParam(r, "status")
	err := usecase.NewUpdateOrderStatusUseCase(oc.Repository).Execute(orderId, orderStatus)

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

func (oc *OrderController) CreateOrderItem(w http.ResponseWriter, r *http.Request) {
	orderId := chi.URLParam(r, "id")
	payload := json.NewDecoder(r.Body)
	orderItem := dto.OrderItemDto{}
	err := payload.Decode(&orderItem)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInvalidRequest(err))
		return
	}

	orderItem.OrderID = orderId
	orderItemSaved, err := usecase.NewCreateOrderItemUseCase(oc.Repository).Execute(orderItem)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	output := &dto.OrderItemResponseDto{
		ID:        orderItemSaved.ID,
		OrderID:   orderItemSaved.OrderID,
		ProductID: orderItemSaved.ProductID,
		Quantity:  orderItemSaved.Quantity,
		Price:     orderItemSaved.Price,
	}

	render.Render(w, r, httpResponsePkg.NewOrderItemResponse(output))
}

func (oc *OrderController) DeleteOrderItem(w http.ResponseWriter, r *http.Request) {
	orderItemId := chi.URLParam(r, "itemId")
	err := usecase.NewDeleteOrderItemUseCase(oc.Repository).Execute(orderItemId)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	render.Render(w, r, httpResponsePkg.NewOrderItemResponse(nil))
}
