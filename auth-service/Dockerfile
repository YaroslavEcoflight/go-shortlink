FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /auth-service ./cmd/server/main.go


FROM alpine:3.21 AS runner

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

COPY --from=builder /auth-service .

EXPOSE 8080

CMD ["./auth-service"]
