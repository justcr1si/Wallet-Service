package postgres

<<<<<<< HEAD
import (
	"context"
	"payment_service/internal/domain"

	"github.com/jackc/pgx/v5"
)

type WalletRepo interface {
	CreateUser(ctx context.Context, email string) (int64, error)
	GetBalance(ctx context.Context, userID int64) (int64, error)
	
	WithTx(ctx context.Context, fn func(tx pgx.Tx) error) error

	GetWalletBalanceForUpdate(ctx context.Context, tx pgx.Tx, userID int64) (int64, error)
	UpdateBalance(ctx context.Context, tx pgx.Tx, userID int64, delta int64) error
	InsertTransaction(ctx context.Context, tx pgx.Tx, t domain.Transaction) error
}
=======
type WalletServiceServer struct {
	
}
>>>>>>> 130ed7e (fix: regenerated code using protoc, feat: created handler implementing gprc server interface)
