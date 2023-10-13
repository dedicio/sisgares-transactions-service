package rabbitmq_utils

import (
	"github.com/streadway/amqp"
)

type Conn struct {
	Conn *amqp.Connection
}

func NewConn(amqpURL string, options ...func(*Conn) error) (*Conn, error) {
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		return nil, err
	}

	brokerConn := &Conn{
		Conn: conn,
	}

	for _, option := range options {
		err := option(brokerConn)
		if err != nil {
			return nil, err
		}
	}

	return brokerConn, nil
}

func (brokerConn *Conn) Close() error {
	return brokerConn.Conn.Close()
}

type Publisher struct {
	Channel *amqp.Channel
}

func NewPublisher(brokerConn *Conn, options ...func(*Publisher) error) (*Publisher, error) {
	channel, err := brokerConn.Conn.Channel()
	if err != nil {
		return nil, err
	}

	brokerChannel := &Publisher{
		Channel: channel,
	}

	for _, option := range options {
		err := option(brokerChannel)
		if err != nil {
			return nil, err
		}
	}

	return brokerChannel, nil
}

func WithPublisherOptionsExchangeName(exchangeName string) func(*Publisher) error {
	return func(brokerChannel *Publisher) error {
		err := brokerChannel.Channel.ExchangeDeclare(
			exchangeName,
			"topic",
			true,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			return err
		}

		return nil
	}
}

func (brokerChannel *Publisher) Close() error {
	return brokerChannel.Channel.Close()
}

func (brokerChannel *Publisher) Publish(
	body []byte,
	routingKey string,
	options ...func(*PublishOptions) error,
) error {
	publishOptions := &PublishOptions{
		ContentType: "application/json",
		Exchange:    "transactions_orders",
	}

	for _, option := range options {
		err := option(publishOptions)
		if err != nil {
			return err
		}
	}

	err := brokerChannel.Channel.Publish(
		publishOptions.Exchange,
		routingKey,
		true,
		false,
		amqp.Publishing{
			ContentType: publishOptions.ContentType,
			Body:        body,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

type PublishOptions struct {
	ContentType string
	Exchange    string
}
