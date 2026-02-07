package domain

import "time"

type User struct {
	ID    int64
	Email string
}

type Transaction struct {
	ID         int
	Type       string
	FromUserID int64
	ToUserID   int64
	Amount     int64
	CreatedAt  time.Time
	Status     string
	Reason     string
}

type CreateUserRequest struct {
	Email string
}

type CreateUserResponse struct {
	User User
}

type GetBalanceRequest struct {
	UserID int64
}

type GetBalanceResponse struct {
	Balance int64
}

type DepositRequest struct {
	UserID int64
	Amount int64
}

type DepositResponse struct {
	Balance int64
}

type WithdrawRequest struct {
	UserID int64
	Amount int64
}

type WithdrawResponse struct {
	Balance int64
}

type TransferRequest struct {
	FromUserID int64
	ToUserID   int64
	Amount     int64
}

type TransferResponse struct {
	FromBalance int64
	ToBalance   int64
}

type ListTransactionsRequest struct {
	UserID int64
	Limit  int64
	Offset int64
}

type ListTransactionsResponse struct {
	Transactions []Transaction
}
