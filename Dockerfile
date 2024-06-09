FROM golang:1.22.3-alpine AS builder

WORKDIR /app

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g main.go

RUN go mod tidy

RUN go build -o main .

FROM alpine:latest as binary

RUN apk add --no-cache tzdata && \
    cp /usr/share/zoneinfo/America/Sao_Paulo /etc/localtime && \
    echo "America/Sao_Paulo" > /etc/timezone && \
    apk del tzdata

COPY --from=builder /app/main .