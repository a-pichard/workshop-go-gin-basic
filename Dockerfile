FROM golang:alpine

WORKDIR /go/src/go-docker-dev.to

COPY . /go/src/go-docker-dev.to

RUN go build .

EXPOSE 8080

CMD ./workshop-go-gin-basic