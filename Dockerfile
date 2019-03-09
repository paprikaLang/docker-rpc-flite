FROM golang as builder

ENV GO111MODULE=on

WORKDIR /server

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY .  .

RUN cd backend && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app


FROM alpine

COPY --from=builder /server/backend/app  /server/

RUN apk update && apk add flite

ENTRYPOINT ["/server/app"]

