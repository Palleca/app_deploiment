# Étape 1 : build
FROM golang:1.25 AS builder

# Répertoire de travail dans le conteneur
WORKDIR /app

# Copier les fichiers go mod pour installer les dépendances
COPY go.mod go.sum ./
RUN go mod download

# Copier tout le projet
COPY . .

# Compiler l’application
RUN go build -o app .

# Étape 2 : image finale (exécution uniquement)
FROM debian:bookworm-slim

WORKDIR /root/

# Copier le binaire compilé depuis l’étape builder
COPY --from=builder /app/app .

# Exposer le port
EXPOSE 8080

# Lancer l’application
CMD ["./app"]
