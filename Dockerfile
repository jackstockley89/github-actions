FROM golang:1.19.2-alpine as build

ARG \
    DIRECTORY \
    COMMAND
ENV \
    CGO_ENABLED=0 \
    GOOS=linux \
    COMMAND=${COMMAND} \
    DIRECTORY=${DIRECTORY}

WORKDIR /go/bin/${DIRECTORY}

RUN apk add --no-cache git

COPY ${DIRECTORY}/lib /go/bin/${DIRECTORY}/lib
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
