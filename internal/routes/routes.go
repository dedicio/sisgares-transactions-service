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

func NewRoutes(db *sql.DB) *Routes {
	return &Routes{
		DB: db,
	}
}

func (routes Routes) Routes() chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.URLFormat)
	router.Use(render.SetContentType(render.ContentTypeJSON))

	orderRepository := repository.NewOrderRepositoryPostgresql(routes.DB)
	orderPublisher := publisher.NewOrderPublisherRabbitmq()

	router.Route("/v1", func(router chi.Router) {
		router.Mount("/orders", NewOrderRoutes(
			orderRepository,
			orderPublisher,
		).Routes())
	})

	return router
}
