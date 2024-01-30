package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aboronilov/go-chi-rest-api/db"
	"github.com/aboronilov/go-chi-rest-api/router"
	"github.com/aboronilov/go-chi-rest-api/services"
	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

type Application struct {
	Config Config
	Models services.Models
}

func (app *Application) Serve() error {
	fmt.Printf("API is running on port %s", app.Config.Port)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", app.Config.Port),
		Handler: router.Routes(),
	}

	return srv.ListenAndServe()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error parsing .env %s", err)
	}

	cfg := Config{
		Port: os.Getenv("PORT"),
	}

	host := os.Getenv("HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable timezone=UTC connect_timeout=5",
		host, port, user, password, dbname,
	)

	dbConn, err := db.ConnectPostgres(dsn)
	if err != nil {
		fmt.Println("No DB connection")
	}
	defer dbConn.DB.Close()

	app := &Application{
		Config: cfg,
		Models: services.New(dbConn.DB),
	}

	err = app.Serve()
	if err != nil {
		log.Fatalf("Error running app %s", err)
	}
}
