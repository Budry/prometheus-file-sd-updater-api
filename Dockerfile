FROM golang:alpine as builder

WORKDIR /go/src/app

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY . .

RUN go install -v ./...

VOLUME /prometheus-file-sd-updater-api


FROM alpine

ENV PORT 80
ENV TOKEN test2

RUN apk add docker

COPY --from=builder /go/bin/prometheus-file-sd-updater-api /usr/bin/prometheus-file-sd-updater-api

CMD prometheus-file-sd-updater-api /prometheus-file-sd-updater-api/targets.json ${TOKEN} --port ${PORT}