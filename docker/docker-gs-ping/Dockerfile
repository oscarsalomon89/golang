# Imagen base oficial de Go
FROM golang:1.23.10-alpine3.22

# Establecer directorio de trabajo
WORKDIR /app

# Copiar archivos del proyecto
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# Compilar la aplicación
RUN go build -o main .

# Exponer el puerto
EXPOSE 8080

# Comando para ejecutar la API
CMD ["./main"]
