# Используем официальный образ Go
FROM golang:1.19-alpine as builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем все файлы проекта в контейнер
COPY . .

# Загружаем зависимости
RUN go mod tidy

# Собираем приложение
RUN go build -o main .

# Запускаем приложение
CMD ["./main"]

