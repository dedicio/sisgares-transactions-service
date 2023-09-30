package routes

import (
	"database/sql"
	"os"

	"github.com/dedicio/sisgares-transactions-service/internal/infra/repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
)

var tokenAuth *jwtauth.JWTAuth

type Routes struct {
	DB *sql.DB
}

func NewRoutes(db *sql.DB) *Routes {
	return &Routes{
		DB: db,
	}
}

func (routes Routes) Routes() chi.Router {
	tokenAuth = jwtauth.New("HS256", []byte(os.Getenv("JWT_SECRET")), nil)
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.URLFormat)
	router.Use(render.SetContentType(render.ContentTypeJSON))
	router.Use(jwtauth.Verifier(tokenAuth))
	router.Use(jwtauth.Authenticator)

	orderRepository := repository.NewOrderRepositoryMysql(routes.DB)

	router.Route("/v1", func(router chi.Router) {
		router.Mount("/orders", NewOrderRoutes(orderRepository).Routes())
	})

	return router
}
