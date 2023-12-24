#build stage
FROM  golang:1.21.5-alpine3.19 AS builder
WORKDIR /app
COPY . .
# Download Go modules
COPY go.mod go.sum ./
RUN go mod download
RUN go build -o main main.go

#run stage
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .
EXPOSE 8000
CMD [ "/app/main" ]