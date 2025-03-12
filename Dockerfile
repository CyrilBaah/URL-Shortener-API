FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o url-shortener

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/url-shortener .
CMD ["./url-shortener"]