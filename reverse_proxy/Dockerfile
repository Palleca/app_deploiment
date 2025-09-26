# Étape 1 : build
FROM golang:1.25 AS builder
WORKDIR /app

COPY go.mod ./
RUN go mod tidy

COPY . .
RUN go build -o proxy .

# Étape 2 : image finale
FROM debian:bookworm-slim
WORKDIR /root/
COPY --from=builder /app/proxy .
EXPOSE 80
CMD ["./proxy"]
