package publisher

import (
	"encoding/json"
	"os"

	"github.com/dedicio/sisgares-transactions-service/internal/entity"
	rabbitmq_utils "github.com/dedicio/sisgares-transactions-service/pkg/rabbitmq"
)

var (
	AMQP_URL = os.Getenv("AMQP_SERVER_URL")
)

type OrderPublisherRabbitmq struct{}

func NewOrderPublisherRabbitmq() *OrderPublisherRabbitmq {
	return &OrderPublisherRabbitmq{}
}

func (ob *OrderPublisherRabbitmq) Create(order *entity.Order) error {
	body, err := json.Marshal(order)
	if err != nil {
		return err
	}
	brokerConn, err := rabbitmq_utils.NewConn(
		AMQP_URL,
	)
	if err != nil {
		return err
	}
	defer brokerConn.Close()

	brokerChannel, err := rabbitmq_utils.NewPublisher(
		brokerConn,
		rabbitmq_utils.WithPublisherOptionsExchangeName("transactions_orders"),
	)
	if err != nil {
		return err
	}
	defer brokerChannel.Close()

	err = brokerChannel.Publish(
		[]byte(body),
		"order.created",
	)
	if err != nil {
		return err
	}

	return nil
}
