package main

import (
	"context"
	"github.com/1makarov/go-dater/internal/config"
	"github.com/1makarov/go-dater/internal/mongodb"
	"github.com/1makarov/go-dater/internal/repository"
	"github.com/1makarov/go-dater/internal/server"
	"github.com/1makarov/go-dater/internal/service"
	"github.com/1makarov/go-dater/pkg/signaler"
	"github.com/1makarov/go-dater/pkg/transport/http"
	"log"
)

//func init() {
//	err := godotenv.Load()
//	if err != nil {
//		log.Fatal("Error loading .env file")
//	}
//}

func main() {
	ctx := context.Background()
	cfg := config.Init()

	client, err := mongodb.Open(ctx, cfg.DB)
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Println(err)
		}
	}()

	db := client.Database(cfg.DB.Name)

	repo := repository.New(db)
	provider := http.New()
	services := service.New(repo, provider)
	servers := server.New(services)
	defer servers.Stop()

	go func() {
		if err = servers.Start(cfg.Port); err != nil {
			log.Println(err)
		}
	}()

	log.Println("started")

	signaler.Wait()
}
