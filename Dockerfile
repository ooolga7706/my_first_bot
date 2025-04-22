# Используем официальный образ Go
FROM golang:1.19-alpine as builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем весь код в контейнер
COPY . .

# Собираем приложение
RUN go mod tidy
RUN go build -o main .

# Запускаем приложение
CMD ["./main"]
