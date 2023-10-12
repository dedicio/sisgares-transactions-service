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
	var items []*entity.OrderItem
	var itemsDto []dto.OrderItemDto

	order := entity.NewOrder(
		items,
		input.Discount,
		input.PaymentMethod,
		input.CompanyId,
	)

	for _, item := range orderItems {
		newItem := entity.NewOrderItem(
			order.ID,
			item.ProductID,
			item.Quantity,
			item.Price,
		)

		items = append(items, &entity.OrderItem{
			ID:        newItem.ID,
			OrderID:   newItem.OrderID,
			ProductID: newItem.ProductID,
			Quantity:  newItem.Quantity,
			Price:     newItem.Price,
		})

		itemsDto = append(itemsDto, dto.OrderItemDto{
			ID:        newItem.ID,
			OrderID:   newItem.OrderID,
			ProductID: newItem.ProductID,
			Quantity:  newItem.Quantity,
			Price:     newItem.Price,
		})
	}
	order.Items = items

	err := uc.Repository.Create(order)
	if err != nil {
		return nil, err
	}

	output := &dto.OrderOutputDto{
		ID:            order.ID,
		Items:         itemsDto,
		Discount:      order.Discount,
		Status:        order.Status,
		PaymentMethod: order.PaymentMethod,
		TotalPrice:    order.TotalPrice(),
		CompanyId:     order.CompanyId,
	}

	return output, nil
}
