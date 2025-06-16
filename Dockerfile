FROM golang:1.24.4-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build .

FROM alpine:latest

RUN apk add --no-cache \
    ca-certificates \
    && update-ca-certificates

WORKDIR /app
COPY --from=builder /app/GoGinInitializer .

ENTRYPOINT ["./GoGinInitializer"]
CMD ["--help"]
