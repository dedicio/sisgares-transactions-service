package usecase

import (
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

func (uc UpdateOrderStatusUseCase) Execute(orderId string, orderStatus string) error {
	err := uc.Repository.UpdateStatus(orderId, orderStatus)
	if err != nil {
		return err
	}

	return nil
}
