package routes

import (
	"github.com/dedicio/sisgares-transactions-service/internal/controllers"
	"github.com/dedicio/sisgares-transactions-service/internal/entity"
	"github.com/go-chi/chi/v5"
)

type OrderRoutes struct {
	Controller controllers.OrderController
}

func NewOrderRoutes(repository entity.OrderRepository) *OrderRoutes {
	return &OrderRoutes{
		Controller: *controllers.NewOrderController(repository),
	}
}

func (or OrderRoutes) Routes() chi.Router {
	router := chi.NewRouter()

	router.Route("/", func(router chi.Router) {
		router.Get("/", or.Controller.FindAll)
		router.Post("/", or.Controller.Create)

		router.Route("/{id}", func(router chi.Router) {
			router.Get("/", or.Controller.FindById)
			router.Patch("/status/{status}", or.Controller.UpdateStatus)
		})
	})

	return router
}
