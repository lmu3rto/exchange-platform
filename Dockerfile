FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN go build -o api ./cmd/api

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/api .

COPY --from=builder /app/docs ./docs 

EXPOSE 8080

CMD ["./api"]