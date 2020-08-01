FROM golang:1.9-alpine

RUN apk update && \
    apk upgrade && \
    apk add --no-cache git

RUN mkdir -p /go/src/app

WORKDIR /go/src/app
RUN go get github.com/golang/dep/cmd/dep
COPY . ./
RUN echo $GOPATH
RUN dep ensure -vendor-only
RUN go build -o main .

CMD "./main"