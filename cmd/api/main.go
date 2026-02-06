package main

import (
	"context"
	"fmt"
	"payment_service/internal/app"
	"payment_service/internal/config"
)

func main() {
	cfg := config.MustLoad()
	ctx := context.Background()

	fmt.Println(cfg)

	app := app.New(ctx, cfg)

	defer app.Repo.Close()

	fmt.Println(app)
}
