# Stage 1: Build binary
FROM golang:1.22 AS builder

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o app .

# Stage 2: Run binary
FROM alpine:3.18

WORKDIR /app
COPY --from=builder /app/app .

ENV PORT=8080
EXPOSE 8080

CMD ["./app"]
