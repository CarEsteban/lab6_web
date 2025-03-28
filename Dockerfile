#usar la imagen oficial de go
FROM golang:1.20-alpine AS builder
WORKDIR /app

#copiar los archivos de modulos y descargar dependencias
COPY go.mod go.sum ./
RUN go mod download

#Copiar el codigo fuente completo
COPY . .

#Compilar el binario
RUN go build -o lab6 ./cmd

#imagen ligear para ejecutar la aplicacino

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/lab6 .
COPY --from=builder /app/LaLigaTracker.html .
EXPOSE 8080
CMD ["./lab6"]
