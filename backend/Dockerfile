FROM golang:1.10.1 as builder

RUN apt-get update && apt-get install -y flite-dev

RUN mkdir /go/src/app

COPY vendor /go/src/vendor

COPY main.go /go/src/app

RUN go install app 


# Removing intermediate container dcf33ff75bf9
FROM debian:jessie-slim 

RUN apt-get update && apt-get install -y flite-dev

COPY --from=builder /go/bin/app /bin/app

ENTRYPOINT ["/bin/app"]