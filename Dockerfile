# Multi-stage Dockerfile para Portfólio Go de Alta Performance

# Etapa 1: Builder
FROM golang:alpine AS builder

ENV GOTOOLCHAIN=auto
WORKDIR /app

# Instala certificados e ferramentas essenciais
RUN apk add --no-cache ca-certificates git

# Cache de dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia código fonte e assets
COPY . .

# Compilação estática sem CGO e com strip de símbolos (-s -w) para tamanho mínimo
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /app/portfolio main.go

# Etapa 2: Runtime Mínimo
FROM alpine:3.20

WORKDIR /app

# Adiciona certificados SSL e curl para healthchecks
RUN apk --no-cache add ca-certificates curl tzdata && \
    addgroup -S appgroup && adduser -S appuser -G appgroup

# Copia apenas o binário compilado que já contém os templates embutidos (embed.FS)
COPY --from=builder /app/portfolio /app/portfolio

# Permissões de usuário seguro não-root
USER appuser

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=5s --start-period=5s --retries=3 \
  CMD curl -f http://localhost:8080/api/health || exit 1

ENTRYPOINT ["/app/portfolio"]
