FROM golang:latest

WORKDIR /payment_service

COPY go.mod .
COPY cmd/payment_service/main.go .

RUN go build -o bin .

ENTRYPOINT [ "/app/bin" ]