FROM golang:1.12.7-alpine3.9

WORKDIR /go/src/app
COPY ./api .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]