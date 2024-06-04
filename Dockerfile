FROM golang:1.22.3-alpine AS builder
WORKDIR /app
COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g cmd/main.go

RUN go build -o main ./cmd

FROM alpine:latest as binary
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]