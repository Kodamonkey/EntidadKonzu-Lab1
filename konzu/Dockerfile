# Usar imagen oficial de Go
FROM golang:1.18

# Crear el directorio de trabajo
WORKDIR /app

# Copiar archivos del servidor Konzu
COPY . .

# Descargar dependencias y compilar el código
RUN go mod tidy
RUN go build -o konzu-app main.go

# Exponer el puerto necesario
EXPOSE 50051

# Comando para ejecutar el servidor
CMD ["./konzu-app"]
