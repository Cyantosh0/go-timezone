FROM golang:alpine
# Required because go requires gcc to build
RUN apk add build-base

RUN apk add --no-cache git

RUN apk add inotify-tools

RUN echo $GOPATH

RUN go install github.com/rubenv/sql-migrate/...@latest

RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY . /timezone_web

WORKDIR /timezone_web

RUN go mod download

CMD sh /timezone_web/docker/run.sh