# ============================================================================
# DOCKER MULTI-STAGE BUILD PARA GO
# ============================================================================
# Multi-stage builds reducen el tamaño de la imagen final
# ============================================================================

# Stage 1: Build
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copiar go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copiar código fuente
COPY . .

# Compilar (CGO disabled para imagen más pequeña)
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/api

# Stage 2: Runtime (imagen mínima)
FROM scratch

# Copiar solo el binario
COPY --from=builder /app/app /app

# Exponer puerto
EXPOSE 8080

# Ejecutar
CMD ["/app"]

