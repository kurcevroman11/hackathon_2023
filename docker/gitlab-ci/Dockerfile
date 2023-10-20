ARG HUB_REGISTRY

FROM golang:1.21.1-alpine3.17

RUN apk --no-cache add netcat-openbsd bash curl

HEALTHCHECK --interval=30s --timeout=10s --retries=3 --start-period=3m CMD curl -m 6 'http://127.0.0.1/graphql' -H 'Accept: application/json' -H 'Content-Type: application/json' --data-binary '{"query":"query HealthCheck { loyaltyServiceStatus { health } }"}' | grep -q "OK" || exit 1

WORKDIR /

EXPOSE 8080

ARG BUILD_FILES_DIR

COPY ../../loyalty/service-loyalty-go.exe /
