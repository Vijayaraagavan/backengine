# syntax=docker/dockerfile:1

FROM golang:1.18-alpine3.15
RUN apk add git

WORKDIR /usr/local/go/src/
RUN go clean -modcache
COPY go.mod ./
COPY go.sum ./

# RUN go mod download

COPY * ./
# RUN go mod download golang.org/x/crypto
RUN go get -u -v -f all
# RUN go mod vendor
RUN go mod tidy

RUN go build -o /backengine

EXPOSE 8080
CMD [ "/backengine" ]
