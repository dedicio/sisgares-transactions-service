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
	_ "github.com/lib/pq"
)

var (
	DB_NAME  = os.Getenv("DB_NAME")
	DB_HOST  = os.Getenv("DB_HOST")
	DB_USER  = os.Getenv("DB_USER")
	DB_PASS  = os.Getenv("DB_PASS")
	DB_PORT  = os.Getenv("DB_PORT")
	AMQP_URL = os.Getenv("AMQP_SERVER_URL")
)

func main() {
	// Database
	dbUrl := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST,
		DB_PORT,
		DB_USER,
		DB_PASS,
		DB_NAME,
	)
	fmt.Println("Connecting to database...", dbUrl)
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		panic(err)
	}

	fmt.Println("Database connection is been established succesfully")
	defer db.Close()

	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	fmt.Println("Message broker connection is been established succesfully")

	router := chi.NewRouter()
	routes := routes.NewRoutes(db)
	router.Use(middleware.Logger)
	router.Mount("/", routes.Routes())

	http.ListenAndServe(":3003", router)
}
