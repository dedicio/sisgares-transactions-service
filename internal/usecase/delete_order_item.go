package usecase

import (
	"github.com/dedicio/sisgares-transactions-service/internal/entity"
)

type DeleteOrderItemUseCase struct {
	Repository entity.OrderRepository
}

func NewDeleteOrderItemUseCase(orderRepository entity.OrderRepository) *DeleteOrderItemUseCase {
	return &DeleteOrderItemUseCase{
		Repository: orderRepository,
	}
}

func (uc DeleteOrderItemUseCase) Execute(orderItemId string) error {
	err := uc.Repository.DeleteOrderItem(orderItemId)
	if err != nil {
		return err
	}

	return nil
}
