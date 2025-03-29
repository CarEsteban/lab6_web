# Usar la imagen oficial de Go (versión 1.20 sobre Alpine Linux) como etapa de construcción.
FROM golang:1.23-alpine AS builder

# Establecer el directorio de trabajo en la imagen.
WORKDIR /app

# Copiar los archivos de módulos (go.mod y go.sum) para aprovechar la caché de dependencias.
COPY go.mod go.sum ./

# Descargar las dependencias especificadas en go.mod.
RUN go mod download

# Copiar el código fuente completo al directorio de trabajo.
COPY . .

# Ejecutar go mod tidy para actualizar go.mod y go.sum según sea necesario.
RUN go mod tidy

# Compilar el binario de la aplicación. Se asume que el entrypoint (main) se encuentra en la carpeta ./cmd.
RUN go build -o lab6 ./cmd

# ---------------------------------------------------------------------
# Etapa final: Imagen ligera para ejecutar la aplicación.
FROM alpine:latest

# Establecer el directorio de trabajo en la imagen final.
WORKDIR /app

# Copiar el binario compilado desde la etapa "builder" al directorio de trabajo.
COPY --from=builder /app/lab6 .

# Copiar el archivo HTML del frontend desde la etapa "builder".
COPY --from=builder /app/LaLigaTracker.html .

# Exponer el puerto 8080 para que la aplicación sea accesible.
EXPOSE 8080

# Comando por defecto para ejecutar la aplicación.
CMD ["./lab6"]

