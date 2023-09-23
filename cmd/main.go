package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/cora23tt/onlinedilerv3"
	"github.com/cora23tt/onlinedilerv3/pkg/handler"
	"github.com/cora23tt/onlinedilerv3/pkg/repository"
	"github.com/cora23tt/onlinedilerv3/pkg/service"
	_ "github.com/lib/pq"
)

func main() {

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "A.Ru.3729#",
		DBName:   "onlineshopv4",
		SSLMode:  "disable",
	})

	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(onlinedilerv3.Server)
	go func() {
		if err := srv.Run("8080", handlers.InitRouts()); err != nil {
			log.Fatalf("error ocured while running http server: %s", err.Error())
		}
	}()

	log.Print("ToDo-app started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print("ToDo-app Shutting Down")
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("error ocured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		log.Printf("error ocured on db connection close: %s", err.Error())
	}

}
