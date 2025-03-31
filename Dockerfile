FROM golang:1.24.1-alpine AS builder

WORKDIR /app

# Устанавливаем часовой пояс
ENV TZ=Europe/Moscow

RUN apk add --no-cache tzdata
RUN ln -sf /usr/share/zoneinfo/Europe/Moscow /etc/localtime

# Копируем файлы go.mod и go.sum и устанавливаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем бинарник
RUN go build -o rest-news ./cmd/main.go

# Финальный образ
FROM alpine:latest

WORKDIR /root/

# Устанавливаем часовой пояс
ENV TZ=Europe/Moscow

RUN apk add --no-cache tzdata && ln -sf /usr/share/zoneinfo/Europe/Moscow /etc/localtime

# Копируем собранное приложение
COPY --from=builder /app/rest-news .


# Запуск
CMD ["./rest-news"]