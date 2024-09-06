# Указываем базовый образ
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy
RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o myapp main.go


# Создаем финальный образ
FROM golang:1.23
WORKDIR /srg

COPY --from=builder /app/myapp .
COPY --from=builder /app/.env .
COPY --from=builder /app/configs.txt .
CMD ["./myapp"]
#CMD ["sh"]
