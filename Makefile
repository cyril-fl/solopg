APP_NAME ?= solopg
VERSION  ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
COMMIT   ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo none)
DATE     ?= $(shell date -u +%Y-%m-%dT%H:%M:%SZ)

LDFLAGS = -X 'solopg/app/cmdrun/buildinfo.AppName=$(APP_NAME)' \
          -X 'solopg/app/cmdrun/buildinfo.Version=$(VERSION)' \
          -X 'solopg/app/cmdrun/buildinfo.Commit=$(COMMIT)' \
          -X 'solopg/app/cmdrun/buildinfo.Date=$(DATE)'

.PHONY: build run

build:
	go build -o $(APP_NAME) -ldflags "$(LDFLAGS)" .

run:
	go run -ldflags "$(LDFLAGS)" .
