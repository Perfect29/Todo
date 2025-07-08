FROM golang:1.24.3-alpine

WORKDIR /app

COPY . .

RUN go build -o main .

CMD ["./main"]