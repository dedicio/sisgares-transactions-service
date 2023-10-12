package usecase

import (
	"github.com/dedicio/sisgares-transactions-service/internal/dto"
	"github.com/dedicio/sisgares-transactions-service/internal/entity"
)

type CreateOrderItemUseCase struct {
	Repository entity.OrderRepository
}

func NewCreateOrderItemUseCase(orderRepository entity.OrderRepository) *CreateOrderItemUseCase {
	return &CreateOrderItemUseCase{
		Repository: orderRepository,
	}
}

func (uc CreateOrderItemUseCase) Execute(input dto.OrderItemDto) (*dto.OrderItemOutputDto, error) {
	orderItem := entity.NewOrderItem(
		input.OrderID,
		input.ProductID,
		input.Quantity,
		input.Price,
	)

	err := uc.Repository.CreateOrderItem(orderItem)
	if err != nil {
		return nil, err
	}

	output := &dto.OrderItemOutputDto{
		ID:        orderItem.ID,
		OrderID:   orderItem.OrderID,
		ProductID: orderItem.ProductID,
		Quantity:  orderItem.Quantity,
		Price:     orderItem.Price,
	}

	return output, nil
}
