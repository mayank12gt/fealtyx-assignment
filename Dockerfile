
FROM golang:1.21 AS builder


WORKDIR /app

COPY go.mod go.sum ./


RUN go mod download


COPY . .


RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/api

FROM alpine:latest


WORKDIR /root/


COPY --from=builder /app/app .


EXPOSE 4000


CMD ["./app"]
