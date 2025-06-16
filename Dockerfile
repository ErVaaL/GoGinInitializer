FROM golang:1.24.4-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o goGinInitializer .

FROM alpine:latest

RUN apk add --no-cache \
    ca-certificates \
    libx11 \
    libxrender \
    libxext \
    libxcursor \
    libxrandr \
    libxi \
    gtk+3.0 \
    && update-ca-certificates

WORKDIR /app
COPY --from=builder /app/goGinInitializer .

ENTRYPOINT ["./goGinInitializer"]
CMD ["--help"]
