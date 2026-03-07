FROM golang:1.25-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /goapp

EXPOSE 8080

CMD ["/goapp"]
