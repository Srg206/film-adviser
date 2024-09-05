# Указываем базовый образ
FROM golang:1.23 AS builder

# Задаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum и загружаем зависимости
COPY go.mod .
COPY go.sum .

 
RUN go mod download

# Копируем остальные файлы приложения
COPY . .

# Компилируем бинарник
RUN  go build -o myapp .

# Создаем финальный образ
FROM alpine:latest

# Копируем бинарник из стадии сборки
COPY --from=builder app/myapp /app/myapp

# Указываем команду для запуска
CMD ["./myapp"]
