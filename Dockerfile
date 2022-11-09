FROM golang:1.19.2-alpine

ARG \
    DIRECTORY \
    COMMAND
ENV \
    CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /go/bin

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY ./${DIRECTORY} .

RUN go build -ldflags "-s -w" .

CMD ["$COMMAND"]