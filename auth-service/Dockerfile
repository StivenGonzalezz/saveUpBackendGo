# Dockerfile
FROM golang:1.24-alpine

# Establece el directorio de trabajo
WORKDIR /saveUp

# Copia los archivos
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# Construye el binario
RUN go build -o main ./cmd/main.go

# Expone el puerto del servicio
EXPOSE 8080

# Ejecuta el binario
CMD ["./main"]
