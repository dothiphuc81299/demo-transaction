FROM golang:alpine

WORKDIR /go/src/app/transaction

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]