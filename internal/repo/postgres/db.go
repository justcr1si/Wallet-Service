package postgres

import (
	"database/sql"
	"fmt"
	"payment_service/internal/config"
)

type Storage struct {
	db *sql.DB
}

func New(cfg *config.Config) (*Storage, error) {
	const op = "storage.postgres.New"

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, 
		cfg.Database.Port, 
		cfg.Database.User, 
		cfg.Database.Password, 
		cfg.Database.Name,
	)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	}

	return &Storage{db: db}, nil

}
