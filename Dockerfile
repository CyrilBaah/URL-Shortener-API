# Build Stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .

# Fix permissions for Go modules
RUN go mod tidy && go build -o url-shortener

# Final Stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/url-shortener .
CMD ["./url-shortener"]
