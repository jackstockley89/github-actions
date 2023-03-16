FROM golang:1.19.2-alpine as build

ARG \
    DIRECTORY \
    COMMAND
ENV \
    CGO_ENABLED=0 \
    GOOS=linux \
    COMMAND=${COMMAND} \
    DIRECTORY=${DIRECTORY}

WORKDIR /go/bin

RUN apk add --no-cache git

COPY ${DIRECTORY}/lib /go/bin/lib
COPY go.mod /go/bin
COPY go.sum /go/bin
RUN go mod download
RUN go install github.com/jackstockley89/github-actions/${DIRECTORY}/lib@latest

COPY ${DIRECTORY}/${COMMAND} /go/bin

RUN go build -ldflags "-s -w" -o ${COMMAND} .

FROM alpine:3.14.2

ARG \
    DIRECTORY \
    COMMAND
ENV \
    COMMAND=${COMMAND}

WORKDIR /go/bin

COPY --from=build /go/bin/${COMMAND} /go/bin

CMD ["sh", "-c", "${COMMAND}"]