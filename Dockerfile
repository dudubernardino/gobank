FROM golang:tip-alpine3.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/api/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .  

EXPOSE 3000

CMD ["./main"]