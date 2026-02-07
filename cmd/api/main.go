package main

import (
	"context"
	"fmt"
	"payment_service/internal/app"
	"payment_service/internal/config"

	log "github.com/sirupsen/logrus"
)

func main() {
	cfg := config.MustLoad()
	ctx := context.Background()

	fmt.Println(cfg)

	application, err := app.New(ctx, cfg)

	if err != nil {
		log.Fatal(err)
	}
	defer application.Close()

	fmt.Println(application)

	if err := application.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
