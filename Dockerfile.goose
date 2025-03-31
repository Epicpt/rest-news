# Dockerfile.goose
FROM golang:1.24.1-alpine

# Устанавливаем необходимые зависимости
RUN apk add --no-cache git

# Устанавливаем goose
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

WORKDIR /migrations