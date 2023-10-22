# Используем базовый образ с Go для сборки
FROM golang:alpine3.18 as builder

COPY . /articles-service

WORKDIR /articles-service

# Установка Git
RUN apk --no-cache add git

COPY ./configs /articles-service/configs

# Копируем файлы проекта и выполняем сборку
COPY ./go.mod ./go.sum ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o articles  github.com/zhashkevych/todo-app/cmd/


# Вторая ступень - образ Alpine для выполнения
FROM alpine:3.14

WORKDIR /articles-service

# Копируем бинарный файл из первой ступени
COPY --from=builder /articles-service/. .

RUN chmod +x /bin/sh && apk --no-cache add nmap bash curl postgresql-client

# Команда для запуска приложения
CMD ["/articles-service/articles"]
