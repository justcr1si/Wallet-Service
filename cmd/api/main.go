package main

import (
	"fmt"
	"payment_service/internal/app"
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
	case envDev:
		log.SetFormatter(&log.JSONFormatter{})
	case envProd:
		log.SetFormatter(&log.JSONFormatter{})
	}
}

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	cfg := config.MustLoad()

	fmt.Println(cfg)

	app := app.New(cfg)
	fmt.Println(app)
}
