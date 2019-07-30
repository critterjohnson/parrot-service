FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o parrot-service .

EXPOSE 3000

CMD ["./parrot-service"]