package main

import (
	"log"
	"os"

	"github.com/Windmill787-golang/junior-test/handler"
	"github.com/Windmill787-golang/junior-test/repository"
	"github.com/Windmill787-golang/junior-test/service"
	"github.com/joho/godotenv"
)

// @title Book API
// @description This is a sample Book CRUD API.
// @version 1.0

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	//repository<-service<-handlers<-routes<-server

	//load environment
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	//init database
	db, err := repository.NewPostgres(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_SSLMODE"),
	)
	if err != nil {
		log.Fatal(err)
	}

	//create repository that depends on database
	r := repository.NewRepository(db)

	//create service that depends on repository
	s := service.NewService(r)

	//create handler that depends on service
	h := handler.NewHandler(s)

	//create and run server that depends on handler routes
	server := NewServer()

	if err = server.Run(os.Getenv("SERVER_PORT"), h.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
