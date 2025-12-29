FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY *.go .

RUN go build -o /app/main -ldflags="-s -w" .

FROM scratch

COPY --from=builder /app/main /main

CMD ["/main"]