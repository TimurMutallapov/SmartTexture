FROM golang:latest

WORKDIR /serve

COPY  . .

RUN go build -o main

CMD ["./main"]
