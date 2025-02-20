FROM golang:1.24-alpine AS builder

WORKDIR /app

# Crear estructura de directorios
RUN mkdir -p services/idp/cmd

# Copiar go.work y todos los go.mod necesarios
COPY backend/go.work* .
COPY backend/services/idp/go.* ./services/idp/
COPY backend/services/tenant-service/go.* ./services/tenant-service/
COPY backend/services/api-gateway/go.* ./services/api-gateway/
COPY backend/shared/common/go.* ./shared/common/
COPY backend/shared/pkg/go.* ./shared/pkg/

# Descargar dependencias
RUN cd services/idp && go mod download

# Copiar el código fuente necesario para el build
COPY backend/services/idp/ ./services/idp/
COPY backend/shared/common/ ./shared/common/
COPY backend/shared/pkg/ ./shared/pkg/

# Compilar la aplicación
RUN CGO_ENABLED=0 GOOS=linux go build -o idp-service ./services/idp/cmd/main.go

# Imagen final
FROM alpine:3.19

WORKDIR /app

# Instalar certificados CA
RUN apk --no-cache add ca-certificates

# Copiar el binario compilado y la configuración
COPY --from=builder /app/idp-service .
COPY backend/services/idp/config/config.yaml ./config/

# Crear usuario no privilegiado
RUN adduser -D appuser
USER appuser

# Exponer el puerto
EXPOSE 8080

CMD ["./idp-service"]