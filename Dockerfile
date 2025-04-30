# Etapa de build
FROM golang:1.22.1
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/app/main.go

# Etapa de execucao
EXPOSE 8080
CMD ["./main"]