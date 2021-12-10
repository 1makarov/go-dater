package main

import (
	"context"
	"github.com/1makarov/go-dater/server/internal/config"
	"github.com/1makarov/go-dater/server/internal/delivery/grpc"
	"github.com/1makarov/go-dater/server/internal/getter"
	"github.com/1makarov/go-dater/server/internal/repository"
	"github.com/1makarov/go-dater/server/internal/server/grpc"
	"github.com/1makarov/go-dater/server/internal/services"
	"github.com/1makarov/go-dater/server/pkg/database/mongodb"
	"github.com/1makarov/go-dater/server/pkg/signaler"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	ctx := context.Background()
	cfg := config.Init()

	clientMongoDB, err := mongodb.DialContext(ctx, mongodb.Input{
		Host:     cfg.DB.Host,
		User:     cfg.DB.User,
		Password: cfg.DB.Password,
		Port:     cfg.DB.Port,
	})
	if err != nil {
		logrus.Errorln(err)
		return
	}
	defer func() {
		if err = clientMongoDB.Disconnect(ctx); err != nil {
			logrus.Errorln(err)
		}
	}()

	db := clientMongoDB.Database(cfg.DB.Name)
	repo := repository.New(db)
	get := getter.New()
	service := services.New(repo, get)
	handler := grpcHandler.NewHandler(service)
	server := grpcServer.New(handler)
	defer server.Stop()

	go func() {
		if err = server.Start(cfg.Port); err != nil {
			logrus.Errorln(err)
			signaler.Signal()
		}
	}()

	logrus.Infoln("started")

	signaler.Wait()
}
