package usecase

import (
	"github.com/dedicio/sisgares-transactions-service/internal/dto"
	"github.com/dedicio/sisgares-transactions-service/internal/entity"
)

type PublishCreateOrderUseCase struct {
	Publisher entity.OrderPublisher
}

func NewPublishCreateOrderUseCase(orderPublisher entity.OrderPublisher) *PublishCreateOrderUseCase {
	return &PublishCreateOrderUseCase{
		Publisher: orderPublisher,
	}
}

func (uc PublishCreateOrderUseCase) Execute(input dto.OrderDto) (*dto.OrderOutputDto, error) {
	orderItems := input.Items
	var items []*entity.OrderItem

	for _, item := range orderItems {
		items = append(items, &entity.OrderItem{
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

	err := uc.Publisher.Create(order)
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
	}

	return output, nil
}
