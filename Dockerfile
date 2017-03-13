FROM golang:1.7-alpine

RUN mkdir -p /go/src/github.com/log-ed/docker-micro

COPY . /go/src/github.com/log-ed/docker-micro

WORKDIR /go/src/github.com/log-ed/docker-micro

RUN go install 

ENTRYPOINT ["/go/bin/docker-micro"]
CMD ["--registry_address=127.0.0.1:8500"]
