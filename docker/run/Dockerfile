# Используем базовый образ с Go для сборки
FROM golang:alpine3.18 as builder

COPY ../../loyalty /loyalty-service

WORKDIR /loyalty-service

# Установка Git
RUN apk --no-cache add git

# Настройка сертификата
RUN apk --no-cache add ca-certificates  && update-ca-certificates

RUN mkdir -p /usr/local/share/ca-certificates/

COPY sebbia.orgRootCA.crt /usr/local/share/ca-certificates/

RUN update-ca-certificates

COPY ../../configs /loyalty-service/configs

# Настройка go.env
RUN go env -w GOPRIVATE="git.apps.okd.sebbia.org/"
RUN go env -w GONOSUMDB="git.apps.okd.sebbia.org/"
RUN go env -w GONOPROXY="git.apps.okd.sebbia.org/"

# Копируем файлы проекта и выполняем сборку
COPY ../../loyalty/go.mod ../../loyalty/go.sum ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o loyalty github.com/loyaltyGo/gqlgen-todos/cmd/



# Вторая ступень - образ Alpine для выполнения
FROM alpine:3.14

WORKDIR /loyalty-service

 # Установка временных зон в Alpine Linux
RUN apk add tzdata

ENV TZ=Europe/Moscow

# Применение новой временной зоны
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime

RUN echo $TZ > /etc/timezone

# Копируем бинарный файл из первой ступени
COPY --from=builder /loyalty-service/. .

RUN chmod +x /bin/sh && apk --no-cache add nmap bash curl postgresql-client

# Команда для запуска приложения
CMD ["/loyalty-service/loyalty"]