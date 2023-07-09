FROM golang:latest

WORKDIR /go/src/github.com/postgres-go

COPY . .

RUN go mod tidy

RUN go mod download

RUN go build -o main .

CMD ["./main"]