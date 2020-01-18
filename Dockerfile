FROM golang:alpine

WORKDIR /go/src/app

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY . .

RUN go install -v ./...

ENTRYPOINT ["prometheus-file-sd-updater-api"]



