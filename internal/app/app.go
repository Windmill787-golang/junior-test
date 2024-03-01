package app

import (
	"log"
	"os"

	"github.com/Windmill787-golang/junior-test/internal/handler"
	"github.com/Windmill787-golang/junior-test/internal/repository"
	"github.com/Windmill787-golang/junior-test/internal/server"
	"github.com/Windmill787-golang/junior-test/internal/service"
	"github.com/joho/godotenv"
)

// Run
// @title Books CRUD API
// @description This is a sample Book CRUD API.
// @version 1.0
// @host localhost:8000
// @BasePath /
// @securityDefinitions.apiKey Bearer
// @in header
// @name Authorization
func Run() {
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
	ser := server.NewServer()

	if err = ser.Run(os.Getenv("SERVER_PORT"), h.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
