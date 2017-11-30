APP?=go-service-skeleton
GOOS?=linux
RELEASE?=0.1

COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

.PHONY: build
build:
	SERVICE_PORT=3000 go build \
		-ldflags "-X main.version=${RELEASE} -X main.commit=${COMMIT} -X main.buildTime=${BUILD_TIME}" \
		-o bin/${GOOS}/${APP}

.PHONY: run
run:
	SERVICE_PORT=3000 go run main.go \
		-ldflags "-X main.version=${RELEASE} -X main.commit=${COMMIT} -X main.buildTime=${BUILD_TIME}"