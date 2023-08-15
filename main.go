package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Windmill787-golang/junior-test/handler"
	"github.com/Windmill787-golang/junior-test/repository"
	"github.com/Windmill787-golang/junior-test/service"
	"github.com/joho/godotenv"
)

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

	//create repository that depents on database
	repository := repository.NewRepository(db)

	//create service that depends on repository
	service := service.NewService(repository)

	//create handler that depends on service
	handler := handler.NewHandler(service)

	fmt.Println(handler)

	//create and run server that depends on handler routes
	server := NewServer()

	//http handler
	handlerFunc := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello default handler")
	}
	handler2 := http.HandlerFunc(handlerFunc)

	http.Handle("/hello", handler2)

	if err = server.Run(os.Getenv("SERVER_PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
