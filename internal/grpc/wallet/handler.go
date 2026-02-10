package handler

import (
	"context"
	pb "payment_service/gen"
	walletv1 "payment_service/gen"
	"payment_service/internal/service/wallet"
)

type Handler struct {
	pb.UnimplementedWalletServiceServer
	service *service.Service
}

func New(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) CreateUser(ctx context.Context, userRequest *walletv1.CreateUserRequest) (*walletv1.CreateUserResponse, error) {
	// заглушка
	return nil, nil
}

func (h *Handler) GetBalance(ctx context.Context, getBalanceRequest *walletv1.GetBalanceRequest) (*walletv1.GetBalanceResponse, error) {
	// заглушка
	return nil, nil
}

func (h *Handler) Deposit(ctx context.Context, depositRequest *walletv1.DepositRequest) (*walletv1.DepositResponse, error) {
	// заглушка
	return nil, nil
}

func (h *Handler) Withdraw(ctx context.Context, withdrawRequest *walletv1.WithdrawRequest) (*walletv1.WithdrawResponse, error) {
	// заглушка
	return nil, nil
}

func (h *Handler) Transfer(ctx context.Context, transferRequest *walletv1.TransferRequest) (*walletv1.TransferResponse, error) {
	// заглушка
	return nil, nil
}

func (h *Handler) ListTransactions(ctx context.Context, listTransactionsRequest *walletv1.ListTransactionsRequest) (*walletv1.ListTransactionsResponse, error) {
	// заглушка
	return nil, nil
}
