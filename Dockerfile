FROM golang:1.14-alpine

WORKDIR /quera
COPY . .

RUN go build -o main .
CMD ["./main"]
