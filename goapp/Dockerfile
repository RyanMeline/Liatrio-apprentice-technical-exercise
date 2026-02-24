FROM golang:1.25-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

EXPOSE 80

RUN CGO_ENABLED=0 GOOS=linux go build -o /goapp

CMD ["/goapp"]
