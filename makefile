DB_URL=postgres://payment:payment@localhost:5432/payment?sslmode=disable

make gen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/wallet/v1/wallet.proto

db-up:
	docker compose up -d

db-down:
	docker compose down

migrate-up:
	migrate -path migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path migrations -database "$(DB_URL)" down 1

migrate-version:
	migrate -path migrations -database "$(DB_URL)" version
