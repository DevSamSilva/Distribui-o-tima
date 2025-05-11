FROM golang:1.22.2

WORKDIR /app

# Primeiro copia apenas os arquivos de módulo (para cache de dependências)
COPY go.mod .
RUN go mod download

# Depois copia o resto e compila
COPY . .
RUN go build -o saque

CMD ["./saque"]