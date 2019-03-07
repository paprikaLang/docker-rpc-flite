FROM golang:1.8

RUN apt-get update && apt-get install -y flite-dev

RUN mkdir /go/src/app

COPY main.go /go/src/app

RUN go install app 

ENTRYPOINT ["/go/bin/app"]