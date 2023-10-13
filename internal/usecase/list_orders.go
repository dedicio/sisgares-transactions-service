package usecase

import (
	"github.com/dedicio/sisgares-transactions-service/internal/dto"
	"github.com/dedicio/sisgares-transactions-service/internal/entity"
)

type ListOrdersUseCase struct {
	Repository entity.OrderRepository
}

func NewListOrdersUseCase(orderRepository entity.OrderRepository) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		Repository: orderRepository,
	}
}

func (uc ListOrdersUseCase) Execute(companyID string) ([]*dto.OrderResponseDto, error) {
	orders, err := uc.Repository.FindAll(companyID)
	if err != nil {
		return nil, err
	}

	var output []*dto.OrderResponseDto
	for _, order := range orders {
		var orderItems []dto.OrderItemDto
		for _, orderItem := range order.Items {
			orderItems = append(orderItems, dto.OrderItemDto{
				ID:        orderItem.ID,
				OrderID:   orderItem.OrderID,
				ProductID: orderItem.ProductID,
				Quantity:  orderItem.Quantity,
				Price:     orderItem.Price,
			})
		}

		output = append(output, &dto.OrderResponseDto{
			ID:            order.ID,
			Discount:      order.Discount,
			Items:         orderItems,
			Status:        order.Status,
			PaymentMethod: order.PaymentMethod,
			TotalPrice:    order.TotalPrice(),
		})
	}

	return output, nil
}
