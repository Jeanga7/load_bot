FROM golang:1.23-alpine AS builder

# Installation des dépendances système (par exemple, git, sqlite, etc.)
# Installation des dépendances système (par exemple, git, sqlite, yt-dlp, etc.)
RUN apk add --no-cache git sqlite gcc musl-dev python3 py3-pip && \
    pip install --no-cache-dir yt-dlp

# Définir le répertoire de travail
WORKDIR /app

# Copier les fichiers go.mod et go.sum, puis télécharger les dépendances
COPY go.mod go.sum ./
RUN go mod download

# Copier le code source dans le conteneur
COPY . .

# Compiler l'application en mode release
RUN CGO_ENABLED=1 GOOS=linux go build -o telegram-bot .

# Image finale, légère
FROM alpine:latest
RUN apk --no-cache add ca-certificates sqlite
WORKDIR /root/
COPY --from=builder /app/telegram-bot .
COPY --from=builder /app/.env .

# Exposer le port pour l'API web (ici 8080)
EXPOSE 8080

# Lancer l'application
CMD ["./telegram-bot"]
