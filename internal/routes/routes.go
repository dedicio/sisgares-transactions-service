package routes

import (
	"database/sql"

	"github.com/dedicio/sisgares-transactions-service/internal/infra/repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type Routes struct {
	DB *sql.DB
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

	router.Route("/v1", func(router chi.Router) {
		router.Mount("/orders", NewOrderRoutes(orderRepository).Routes())
	})

	return router
}
