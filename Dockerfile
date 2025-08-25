# Stage 1: Build the application
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -ldflags="-s -w" -o main ./cmd/server/main.go

# Stage 2: Runtime
FROM alpine:3.20

WORKDIR /app

# Install netcat and postgresql-client for pg_isready
RUN apk add --no-cache netcat-openbsd postgresql-client

COPY --from=builder /app/main .
COPY docker-entrypoint.sh .

RUN chmod +x docker-entrypoint.sh

EXPOSE 8080

ENTRYPOINT ["./docker-entrypoint.sh"]