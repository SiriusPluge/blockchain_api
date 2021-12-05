
FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o main ./cmd/

CMD ["./main"]