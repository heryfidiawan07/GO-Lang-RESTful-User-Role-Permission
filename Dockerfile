# Stage 1: Build binary
FROM golang:1.22 AS builder

WORKDIR /app
COPY . .

RUN go mod tidy
# 👇 compile statis, tidak pakai libc/glibc
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .

# Stage 2: Run binary
FROM scratch

WORKDIR /app
COPY --from=builder /app/app .

ENV PORT=8080
EXPOSE 8080

CMD ["./app"]
