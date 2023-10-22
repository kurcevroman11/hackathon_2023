FROM postgres:10.5-alpine

RUN mkdir -p /Docker-entrypoint-initdb.d
COPY ./init.sh /docker-entrypoint-initdb.d/init.sh
RUN apk add --no-cache bash