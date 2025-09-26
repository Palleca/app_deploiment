# Étape 1 : build et tests
FROM golang:1.25 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# run tests avec couverture
RUN go test ./... -coverprofile=coverage.out

# build du binaire
RUN go build -o app .

# Étape 2 : image finale
FROM debian:bookworm-slim
WORKDIR /root/
COPY --from=builder /app/app .
EXPOSE 8080
CMD ["./app"]
