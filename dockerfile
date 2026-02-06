FROM golang:latest

WORKDIR /wallet-service

COPY go.mod .
COPY cmd/payment_service/main.go .

RUN go build -o bin .

ENTRYPOINT [ "/cmd/api" ]