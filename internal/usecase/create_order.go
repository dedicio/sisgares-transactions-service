package usecase

import (
	"github.com/dedicio/sisgares-transactions-service/internal/dto"
	"github.com/dedicio/sisgares-transactions-service/internal/entity"
)

type CreateOrderUseCase struct {
	Repository entity.OrderRepository
}

func NewCreateOrderUseCase(orderRepository entity.OrderRepository) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		Repository: orderRepository,
	}
}

func (uc CreateOrderUseCase) Execute(input dto.OrderDto) (*dto.OrderOutputDto, error) {
	orderItems := input.Items
	var items []entity.OrderItem

	for _, item := range orderItems {
		items = append(items, entity.OrderItem{
			ID:        item.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
		})
	}

	order := entity.NewOrder(
		items,
		input.Discount,
		input.PaymentMethod,
		input.CompanyId,
	)

	err := uc.Repository.Create(order)
	if err != nil {
		return nil, err
	}

	output := &dto.OrderOutputDto{
		ID:         order.ID,
		Items:      orderItems,
		Discount:   order.Discount,
		Status:     order.Status,
		TotalPrice: order.TotalPrice(),
		CompanyId:  order.CompanyId,
		CreatedAt:  order.CreatedAt,
		UpdatedAt:  order.UpdatedAt,
	}

	return output, nil
}
