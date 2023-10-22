# Используем базовый образ с Go для сборки
FROM golang:alpine3.18 as builder

COPY ../../. /article-service

WORKDIR /article-service


# Копируем файлы проекта и выполняем сборку
COPY ../.././go.mod ../.././go.sum ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o article github.com/zhashkevych/todo-app/cmd/

# Вторая ступень - образ Alpine для выполнения
FROM alpine:3.14

WORKDIR /article-service

# Копируем бинарный файл из первой ступени
COPY --from=builder /article-service/. .

RUN chmod +x /bin/sh && apk --no-cache add nmap bash curl postgresql-client

# Команда для запуска приложения
CMD ["/article-service/article"]