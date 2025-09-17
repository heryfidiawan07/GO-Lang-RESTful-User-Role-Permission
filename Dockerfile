# Stage 1: build binary
FROM golang:1.22 AS builder

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o app .

# Stage 2: run binary
FROM debian:bullseye-slim

WORKDIR /app
COPY --from=builder /app/app .

# Railway akan inject PORT env
ENV PORT=8080
EXPOSE 8080

CMD ["./app"]
