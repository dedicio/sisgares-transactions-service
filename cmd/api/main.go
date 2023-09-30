package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dedicio/sisgares-transactions-service/internal/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DB_NAME = os.Getenv("DB_NAME")
	DB_HOST = os.Getenv("DB_HOST")
	DB_USER = os.Getenv("DB_USER")
	DB_PASS = os.Getenv("DB_PASS")
	DB_PORT = os.Getenv("DB_PORT")
)

func main() {
	dbUrl := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		DB_USER,
		DB_PASS,
		DB_HOST,
		DB_PORT,
		DB_NAME,
	)
	db, err := sql.Open("mysql", dbUrl)

	if err != nil {
		panic(err)
	}

	fmt.Println("Database connection is been established succesfully")
	defer db.Close()

	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	router := chi.NewRouter()
	routes := routes.NewRoutes(db)
	router.Use(middleware.Logger)
	router.Mount("/", routes.Routes())

	http.ListenAndServe(":3000", router)
}

// func apiVersionCtx(version string) func(next http.Handler) http.Handler {
// 	return func(next http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			r = r.WithContext(context.WithValue(r.Context(), "api.version", version))
// 			next.ServeHTTP(w, r)
// 		})
// 	}
// }
