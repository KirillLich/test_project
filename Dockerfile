FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY cmd/main.go .

RUN go build -o /app/main -ldflags="-s -w" .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main /main

EXPOSE 8080

CMD ["/main"]