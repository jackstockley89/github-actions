FROM golang:1.19.2-alpine as build

ARG \
    DIRECTORY \
    COMMAND
ENV \
    CGO_ENABLED=0 \
    GOOS=linux \
    COMMAND=${COMMAND} \
    DIRECTORY=${DIRECTORY}

WORKDIR /go/bin/lib

RUN apk add --no-cache git

COPY lib/go.mod /go/bin/lib
COPY lib/go.sum /go/bin/lib
RUN go mod download

RUN go install github.com/jackstockley89/github-actions/lib/github@api

COPY lib/${DIRECTORY} /go/bin/lib/${DIRECTORY}

WORKDIR /go/bin/${DIRECTORY}

COPY ${DIRECTORY}/go.mod /go/bin/${DIRECTORY}
COPY ${DIRECTORY}/go.sum /go/bin/${DIRECTORY}
RUN go mod download

COPY ${DIRECTORY}/${COMMAND} /go/bin/${DIRECTORY}

RUN go build -ldflags "-s -w" -o ${COMMAND} .

FROM golang:1.19.2-alpine

ARG \
    DIRECTORY \
    COMMAND
ENV \
    COMMAND=${COMMAND}

WORKDIR /go/bin

COPY --from=build /go/bin/${DIRECTORY}/${COMMAND} /go/bin

CMD ["sh", "-c", "${COMMAND}"]
