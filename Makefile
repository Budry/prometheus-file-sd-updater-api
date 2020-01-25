OUTPUT=prometheus-file-sd-updater-api

all: build
build:
		docker build -t budry/prometheus-file-sd-updater-api .
		docker push budry/prometheus-file-sd-updater-api