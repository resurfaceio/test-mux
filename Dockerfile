FROM golang:1.16

WORKDIR /app

COPY /src .

RUN go mod download

RUN go build .

CMD ["./test-mux"]