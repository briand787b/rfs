FROM golang:1.11

RUN apt-get update && apt-get upgrade -y

WORKDIR /go/src/github.com/briand787b/rfs

EXPOSE 8080

COPY . .

RUN go install

ENTRYPOINT [ "rfs" ]