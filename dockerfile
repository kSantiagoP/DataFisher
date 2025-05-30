#Go Images definition
FROM golang:1.24.2-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o /app/dataFisher-api ./cmd/api
RUN go build -o /app/dataFisher-worker ./cmd/worker
RUN go build -o /app/dataFisher-db-init ./cmd/db_init

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/.env .
COPY --from=builder /app/dataFisher-api .
COPY --from=builder /app/dataFisher-worker .
COPY --from=builder /app/dataFisher-db-init .
COPY --from=builder /app/internal/data_api/mock ./internal/data_api/mock
CMD ["./dataFisher-api"]