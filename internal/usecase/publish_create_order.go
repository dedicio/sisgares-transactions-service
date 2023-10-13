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

func (uc PublishCreateOrderUseCase) Execute(input *dto.OrderOutputDto) error {
	var items []*entity.OrderItem
	for _, item := range input.Items {
		newItem := entity.NewOrderItem(
			input.ID,
			item.ProductID,
			item.Quantity,
			item.Price,
		)

		items = append(items, newItem)
	}

	order := entity.NewOrder(
		items,
		input.Discount,
		input.PaymentMethod,
		input.CompanyId,
	)

	err := uc.Publisher.Create(order)
	if err != nil {
		return err
	}

	return nil
}
