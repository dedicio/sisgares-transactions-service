module github.com/dedicio/sisgares-transactions-service

go 1.21.0

require (
	github.com/go-chi/chi/v5 v5.2.2
	github.com/go-chi/render v1.0.3
	github.com/google/uuid v1.3.1
)

require github.com/rabbitmq/amqp091-go v1.7.0 // indirect

require (
	github.com/ajg/form v1.5.1 // indirect
	github.com/lib/pq v1.10.9
	github.com/streadway/amqp v1.1.0
	github.com/wagslane/go-rabbitmq v0.12.4
)
