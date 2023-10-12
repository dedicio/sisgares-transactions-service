package routes

import (
	"database/sql"

	"github.com/dedicio/sisgares-transactions-service/internal/infra/publisher"
	"github.com/dedicio/sisgares-transactions-service/internal/infra/repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/wagslane/go-rabbitmq"
)

type Routes struct {
	DB         *sql.DB
	BrokerConn *rabbitmq.Conn
}

func NewRoutes(db *sql.DB, brokerConn *rabbitmq.Conn) *Routes {
	return &Routes{
		DB:         db,
		BrokerConn: brokerConn,
	}
}

func (routes Routes) Routes() chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.URLFormat)
	router.Use(render.SetContentType(render.ContentTypeJSON))

	orderRepository := repository.NewOrderRepositoryPostgresql(routes.DB)

	brokerChannel, err := createPublisher(routes.BrokerConn)
	if err != nil {
		panic(err)
	}
	defer brokerChannel.Close()

	publisher := publisher.NewOrderPublisherRabbitmq(brokerChannel)

	router.Route("/v1", func(router chi.Router) {
		router.Mount("/orders", NewOrderRoutes(
			orderRepository,
			publisher,
		).Routes())
	})

	return router
}

func createPublisher(conn *rabbitmq.Conn) (*rabbitmq.Publisher, error) {
	return rabbitmq.NewPublisher(
		conn,
		rabbitmq.WithPublisherOptionsLogging,
		rabbitmq.WithPublisherOptionsExchangeName("transactions_orders"),
		rabbitmq.WithPublisherOptionsExchangeDeclare,
	)
}
