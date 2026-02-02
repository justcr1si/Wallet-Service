package wallet

import (
	"payment_service/internal/config"
	"payment_service/internal/repo/postgres"

	log "github.com/sirupsen/logrus"
)

type Service struct {
	storage *postgres.Storage
}

func New(cfg *config.Config) (*Service, error) {
	storage, err := postgres.New(cfg)

	if err != nil {
		log.Errorf("Error: %v", err)
	}

	return &Service{
		storage: storage,
	}, nil
}
