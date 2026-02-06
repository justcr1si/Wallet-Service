package observability

import (
	"payment_service/internal/config"

	log "github.com/sirupsen/logrus"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func init() {
	cfg := config.MustLoad()

	switch cfg.Env {
	case envLocal:
		log.SetFormatter(&log.TextFormatter{})
	case envDev, envProd:
		log.SetFormatter(&log.JSONFormatter{})
	}
}
