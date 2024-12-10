FROM golang:1.23-alpine

WORKDIR /app

# Copy semua file ke dalam container
COPY . .

# Unduh dependensi Go
RUN go mod download

# Build aplikasi
RUN go build -o main .

# Expose port 8080 (port default untuk Cloud Run)
EXPOSE 8080

# Set environment variable untuk PORT (opsional, jika aplikasi Go Anda membutuhkan pengaturan ini)
ENV PORT=8080

# Jalankan aplikasi
CMD ["./main"]
