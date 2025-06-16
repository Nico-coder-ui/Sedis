FROM golang:1.24.1 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o sedis_manager .

FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder /app/sedis_manager .

EXPOSE 8087
CMD ["./sedis_manager"]
