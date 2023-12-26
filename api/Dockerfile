FROM golang:1.21.5

WORKDIR /go/api/
COPY . .

RUN apt-get update && apt-get install -y git vim

RUN go install && go build -o /go/bin/air github.com/cosmtrek/air
