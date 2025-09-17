# Stage 1: Build binary
FROM golang:1.22 AS builder

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o app .

# Stage 2: Run binary
FROM debian:bullseye-slim

WORKDIR /app
COPY --from=builder /app/app .

# Railway inject $PORT otomatis
ENV PORT=8080
EXPOSE 8080

CMD ["./app"]
