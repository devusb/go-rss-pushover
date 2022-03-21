FROM golang:1.15-alpine
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /go-rss-pushover

CMD [ "/go-rss-pushover" ]