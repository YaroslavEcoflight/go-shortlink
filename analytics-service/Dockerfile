FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /analytics-service ./cmd/server/main.go

FROM alpine:3.21 AS runner

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

COPY --from=builder /analytics-service .

EXPOSE 8082

CMD ["./analytics-service"]
