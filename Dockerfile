# Gunakan image Golang resmi sebagai base image
FROM golang:latest as builder

# Set environment variable agar Go menggunakan mode production
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Buat direktori kerja di dalam container
WORKDIR /app

# Salin file go.mod dan go.sum terlebih dahulu dan lakukan download dependensi
COPY go.mod .
COPY go.sum .
RUN go mod download

# Salin seluruh kode sumber aplikasi
COPY . .
COPY .env .

# Build aplikasi Golang
RUN go build -o be-shop ./cmd/be-shop

# Expose port yang digunakan oleh aplikasi
EXPOSE 8060

# Atur command untuk dijalankan saat container dijalankan
CMD ["go", "run", "cmd/be-shop/main.go"]
