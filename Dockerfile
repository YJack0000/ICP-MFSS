FROM golang:1.21.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

FROM ubuntu:22.04

WORKDIR /app

COPY --from=builder /app/static ./static
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
