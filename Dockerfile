# Etapa de build
FROM golang:1.23 AS builder
WORKDIR /app

# Copia os arquivos de dependência do Go
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante do código para o contêiner
COPY . .

# Compila o binário
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o ordersystem ./cmd/ordersystem

# Etapa final
FROM alpine:latest
WORKDIR /app

# Copia o binário da etapa de build
COPY --from=builder /app/ordersystem .

# Copia o arquivo .env, se necessário
COPY .env .env

# Define a porta exposta (ajuste conforme necessário)
EXPOSE 8080

# Comando para rodar a aplicação
CMD ["/app/ordersystem"]