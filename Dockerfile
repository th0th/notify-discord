FROM golang:1.17

WORKDIR /usr/src/notify-discord

COPY go.mod go.sum ./
RUN go mod download

COPY config.go .
COPY description_part.go .
COPY main.go .
COPY validator.go .

RUN mkdir bin

RUN CGO_ENABLED=0 go build -a -o bin/notify-discord .

FROM alpine:latest

RUN apk update && apk add bash

COPY --from=0 /usr/src/rancher-redeploy-workload/bin/notify-discord /usr/local/bin/notify-discord

COPY docker-entrypoint.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/docker-entrypoint.sh

CMD ["/usr/local/bin/docker-entrypoint.sh"]
