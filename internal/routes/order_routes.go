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

func (lr OrderRoutes) Routes() chi.Router {
	router := chi.NewRouter()

	router.Route("/", func(router chi.Router) {
		router.Get("/", lr.Controller.FindAll)
		router.Post("/", lr.Controller.Create)

		router.Route("/{id}", func(router chi.Router) {
			router.Get("/", lr.Controller.FindById)
			router.Patch("/status/{status}", lr.Controller.UpdateStatus)
		})
	})

	return router
}
