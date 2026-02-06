package wallet

import (
	"context"
	"payment_service/internal/config"
	"payment_service/internal/repo/postgres"

	log "github.com/sirupsen/logrus"
)

type Service struct {
	Storage *postgres.Storage
}

func New(ctx context.Context, cfg *config.Config) (*Service, error) {
	storage, err := postgres.New(ctx, cfg)

	if err != nil {
		log.Errorf("Error: %v", err)
	}

	return &Service{
		Storage: storage,
	}, nil
}
