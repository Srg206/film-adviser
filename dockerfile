# Указываем базовый образ
FROM golang:1.23 AS builder

# Задаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum и загружаем зависимости
COPY go.mod go.sum ./
RUN go mod tidy
RUN go mod download

# Копируем остальные файлы приложения
COPY . .

# Компилируем бинарник
RUN  go build -o myapp main.go



# Создаем финальный образ
FROM alpine:latest
WORKDIR /app
# Копируем бинарник из стадии сборки
COPY --from=builder app/myapp .

# Указываем команду для запуска
CMD ["./myapp"]