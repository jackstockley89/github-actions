FROM golang:1.19.2-alpine

ARG \
    DIRECTORY \
    COMMAND
ENV \
    CGO_ENABLED=0 \
    GOOS=linux \
    COMMAND=${COMMAND} \
    DIRECTORY=${DIRECTORY}

WORKDIR /go/bin

COPY ./github-api /go/bin/github-api
COPY go.mod /go/bin
COPY go.sum /go/bin
RUN go mod download
COPY ./${DIRECTORY} /go/bin

RUN go build -ldflags "-s -w" -o ${COMMAND} ./${DIRECTORY}

CMD ["sh", "-c", "${COMMAND}"]