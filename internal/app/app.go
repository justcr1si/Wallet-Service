package app

import (
	"payment_service/internal/config"
	"payment_service/internal/service/wallet"

	log "github.com/sirupsen/logrus"
)

type App struct {
	cfg     *config.Config
	service *wallet.Service
}

func New(cfg *config.Config) *App {
	service, err := wallet.New(cfg)

	if err != nil {
		log.Errorf("Error: %v", err)
	}

	return &App{
		cfg:     cfg,
		service: service,
	}
}
