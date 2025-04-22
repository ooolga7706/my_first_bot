# Используем официальный образ Go
FROM golang:1.19-alpine as builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы Go
COPY . .

# Загружаем зависимости
RUN go mod tidy

# Собираем программу
RUN go build -o main .

# Запускаем бота
CMD ["./main"]

