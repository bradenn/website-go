FROM golang:1.16

WORKDIR /go/src/bradenn
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["bradenn"]