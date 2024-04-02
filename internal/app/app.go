package app

import (
	"fmt"
	"log"

	"github.com/Windmill787-golang/junior-test/internal/config"
	"github.com/Windmill787-golang/junior-test/internal/handler"
	"github.com/Windmill787-golang/junior-test/internal/repository"
	"github.com/Windmill787-golang/junior-test/internal/server"
	"github.com/Windmill787-golang/junior-test/internal/service"

	"github.com/joho/godotenv"
)

const (
	ConfigDir  = "configs"
	ConfigFile = "main"
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

	//load config
	c, err := config.New(ConfigDir, ConfigFile)
	if err != nil {
		log.Fatal(err)
	}

	//init database
	db, err := repository.NewPostgres(&c.Postgres)
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
	ser := server.NewServer(h, &c.Server)

	fmt.Println("Server started")

	//run server
	if err = ser.Run(); err != nil {
		log.Fatal(err)
	}
}
