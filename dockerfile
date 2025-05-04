#Go Images definition
FROM golang:1.24.2-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o /app/dataFisher-app ./cmd/api

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/dataFisher-app .
COPY --from=builder /app/internal/data_api/mock ./internal/data_api/mock
CMD ["./dataFisher-app"]