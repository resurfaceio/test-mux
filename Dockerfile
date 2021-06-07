FROM golang:1.16

WORKDIR /src

COPY . .
RUN ls 

RUN go mod download

RUN ls .
RUN go build .

EXPOSE 5000
CMD [/src/main]
