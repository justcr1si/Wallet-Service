package app

import (
	"context"
	"payment_service/internal/config"
	"payment_service/internal/repo/postgres"
	"payment_service/internal/service/wallet"

	log "github.com/sirupsen/logrus"
)

type App struct {
	Cfg     *config.Config
	Service *wallet.Service
	Repo    *postgres.Storage
}

func New(ctx context.Context, cfg *config.Config) *App {
	service, err := wallet.New(ctx, cfg)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	repo, err := postgres.New(ctx, cfg)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	return &App{
		Cfg:     cfg,
		Service: service,
		Repo:    repo,
	}
}
