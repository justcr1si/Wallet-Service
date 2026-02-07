package wallet_handler

import "payment_service/internal/service/wallet"

type WalletHandler struct {
	Service *wallet.Service
}

func NewWalletHandler(service *wallet.Service) (*WalletHandler, error) {
	return &WalletHandler{
		Service: service,
	}, nil
}
