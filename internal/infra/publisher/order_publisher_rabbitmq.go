package publisher

import (
	"github.com/dedicio/sisgares-transactions-service/internal/entity"
	rabbitmq "github.com/wagslane/go-rabbitmq"
)

type OrderPublisherRabbitmq struct {
	BrokenChannel *rabbitmq.Publisher
}

func NewOrderPublisherRabbitmq(brokerChannel *rabbitmq.Publisher) *OrderPublisherRabbitmq {
	return &OrderPublisherRabbitmq{
		BrokenChannel: brokerChannel,
	}
}

func (ob *OrderPublisherRabbitmq) Create(order *entity.Order) error {
	err := ob.BrokenChannel.Publish(
		[]byte(order.ID),
		[]string{"transactions_orders_create"},
		rabbitmq.WithPublishOptionsContentType("application/json"),
		rabbitmq.WithPublishOptionsExchange("transactions_orders"),
	)
	if err != nil {
		return err
	}

	return nil
}
