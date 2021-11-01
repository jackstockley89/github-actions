FROM golang:1.17-alpine3.13

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build .

CMD ["auth-check-approve"]