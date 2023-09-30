package usecase

import (
	"github.com/dedicio/sisgares-transactions-service/internal/dto"
	"github.com/dedicio/sisgares-transactions-service/internal/entity"
)

type UpdateOrderStatusUseCase struct {
	Repository entity.OrderRepository
}

func NewUpdateOrderStatusUseCase(orderRepository entity.OrderRepository) *UpdateOrderStatusUseCase {
	return &UpdateOrderStatusUseCase{
		Repository: orderRepository,
	}
}

func (uc UpdateOrderStatusUseCase) Execute(input dto.OrderDto) error {
	err := uc.Repository.UpdateStatus(input.ID, input.Status)
	if err != nil {
		return err
	}

	return nil
}
