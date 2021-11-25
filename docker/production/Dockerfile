# Build stage
FROM golang:1.17.3-alpine3.14 AS builder
WORKDIR /app
COPY . .
RUN go build -o plant-system-api main.go

# Run stage
FROM alpine:3.15
WORKDIR /app
COPY --from=builder /app/plant-system-api .

EXPOSE 8080
CMD [ "/app/plant-system-api" ]