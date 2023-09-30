package usecase

import (
	"github.com/dedicio/sisgares-transactions-service/internal/dto"
	"github.com/dedicio/sisgares-transactions-service/internal/entity"
)

type UpdateOrderUseCase struct {
	Repository entity.OrderRepository
}

func NewUpdateOrderUseCase(orderRepository entity.OrderRepository) *UpdateOrderUseCase {
	return &UpdateOrderUseCase{
		Repository: orderRepository,
	}
}

func (uc UpdateOrderUseCase) Execute(input dto.OrderDto) error {
	err := uc.Repository.UpdateStatus(input.ID, input.Status)
	if err != nil {
		return err
	}

	return nil
}
