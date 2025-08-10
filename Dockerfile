FROM golang:1.21-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /bin/telephony-gateway ./cmd/server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /bin/telephony-gateway /app/telephony-gateway

ENV TELEPHONY_GATEWAY_SERVICE_PORT=8080

CMD ["/app/telephony-gateway"]
