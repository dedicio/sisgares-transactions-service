package usecase

import (
	"github.com/dedicio/sisgares-transactions-service/internal/dto"
	"github.com/dedicio/sisgares-transactions-service/internal/entity"
)

type FindOrderByIdUseCase struct {
	OrderRepository entity.OrderRepository
}

func NewFindOrderByIdUseCase(orderRepository entity.OrderRepository) *FindOrderByIdUseCase {
	return &FindOrderByIdUseCase{
		OrderRepository: orderRepository,
	}
}

func (uc FindOrderByIdUseCase) Execute(id string) (*dto.OrderResponseDto, error) {
	order, err := uc.OrderRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	items, err := uc.OrderRepository.FindAllOrderItemsByOrderId(id)
	if err != nil {
		return nil, err
	}

	var orderItems []dto.OrderItemDto
	for _, orderItem := range items {
		orderItemDto := dto.OrderItemDto{
			ID:        orderItem.ID,
			ProductID: orderItem.ProductID,
			Quantity:  orderItem.Quantity,
			Price:     orderItem.Price,
		}

		orderItems = append(orderItems, orderItemDto)
	}

	return &dto.OrderResponseDto{
		ID:            order.ID,
		Items:         orderItems,
		Discount:      order.Discount,
		Status:        order.Status,
		PaymentMethod: order.PaymentMethod,
		TotalPrice:    order.TotalPrice(),
	}, nil
}
