OUTPUT=prometheus-file-sd-updater-api

all: build
build:
		go build -o dist/$(OUTPUT) -v