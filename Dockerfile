# Etapa 1: Construcción del binario
FROM golang:1.20 AS builder

# Establece el directorio de trabajo
WORKDIR /app

# Copia los archivos del módulo y descarga las dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copia el código fuente
COPY . .

# Compila el binario
RUN go build -o main .

# Compila el binario con las opciones necesarias para evitar dependencias externas
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Etapa 2: Imagen final
FROM alpine:latest

# Instala dependencias necesarias con apk *(Gestor de paquetes de imagen Alpine).
# --no-cache evita que los archivos temporales se mantengan en el contenedor.
RUN apk --no-cache add ca-certificates

# Copia el binario desde la etapa de construcción
COPY --from=builder /app/main /app/main

# Define el comando por defecto para ejecutar el binario
ENTRYPOINT ["/app/main"]

# Expone el puerto en el que tu aplicación escucha
EXPOSE 90
