package main

import (
	"log"

	"github.com/lafetz/url-shortner/internal/adapters/primary/web"
	"github.com/lafetz/url-shortner/internal/adapters/secondary/repository"
	"github.com/lafetz/url-shortner/internal/core/services"
)

func main() {
	db, err := repository.OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.NewDb(db)
	userService := services.NewUserService(repo)
	urlService := services.NewUrlService(repo)
	application := web.NewApp(userService, *urlService)
	application.Run()
}
