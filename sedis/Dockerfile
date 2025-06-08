FROM golang:1.24.1 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o sedis .

FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder /app/sedis .

EXPOSE 8085
EXPOSE 8086
CMD ["./sedis"]
