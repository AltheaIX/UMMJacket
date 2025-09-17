FROM golang:1.23.6-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go run github.com/swaggo/swag/cmd/swag init --generalInfo ./internal/cmd/main.go
RUN go build -o main ./internal/cmd/

CMD ["./main"]