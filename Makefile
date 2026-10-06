ENV ?= dev

APP_NAME ?= solopg
VERSION  ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
COMMIT   ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo none)
DATE     ?= $(shell date -u +%Y-%m-%dT%H:%M:%SZ)

LDFLAGS = -X solopg/app/cmdversion/services/process/generatebuildinfo.AppName=$(APP_NAME) \
          -X solopg/app/cmdversion/services/process/generatebuildinfo.Version=$(VERSION) \
          -X solopg/app/cmdversion/services/process/generatebuildinfo.Commit=$(COMMIT) \
          -X solopg/app/cmdversion/services/process/generatebuildinfo.Date=$(DATE) \
          -X solopg/config.env=$(ENV)

.PHONY: build run

build:
	go build -o $(APP_NAME) -ldflags "$(LDFLAGS)" .

run:
	go run -ldflags "$(LDFLAGS)" .
