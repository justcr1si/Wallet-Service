package app

import (
	"context"
	// "net"
	"payment_service/internal/config"
	// wallet_handler "payment_service/internal/grpc/wallet"
	// pb "payment_service/gen/wallet/v1"
	"payment_service/internal/repo/postgres"
	"payment_service/internal/service/wallet"

	// log "github.com/sirupsen/logrus"
	// "google.golang.org/grpc"
)

type App struct {
	Cfg     *config.Config
	Service *wallet.Service
	Repo    *postgres.Storage
}

func New(ctx context.Context, cfg *config.Config) (*App, error) {
	service, err := wallet.New(ctx, cfg)

	if err != nil {
		return nil, err
	}

	repo, err := postgres.New(ctx, cfg)

	if err != nil {
		repo.Close()
		return nil, err
	}

	return &App{
		Cfg:     cfg,
		Service: service,
		Repo:    repo,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	// log.Info("App started")

	// // заглушка
	// // <-ctx.Done()
	// lis, err := net.Listen("tcp", a.Cfg.GRPC.Port)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// grpcServer := grpc.NewServer()
	// handler, err := wallet_handler.NewWalletHandler(a.Service)
	// pb.RegisterWalletServiceServer(grpcServer, handler)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Info("App stopped")

	// go func() {
	// 	<-ctx.Done()
	// 	grpcServer.GracefulStop()
	// }()

	// log.Infof("gRPC listening on %s", a.Cfg.GRPC.Port)
	// return grpcServer.Serve(lis)
	return nil
}

func (a *App) Close() {
	if a.Repo != nil {
		a.Repo.Close()
	}
}
