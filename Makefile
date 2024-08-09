SHELL := /bin/sh
.DEFAULT_GOAL := build

VENDOR := nikko

override VERSION := v0.0.1
override TARGET := clickhouse_service
override MODE := prod

MAIN := ./app/cmd/$(TARGET)/main.go
IMAGENAME := $(VENDOR)/$(shell echo $(TARGET) | tr A-Z a-z)

#PARAMS :=

# Operating System Default (LINUX)
TARGETOS=linux

# Flags for build
LDFLAGS = "-s -w -X main.Version=$(VERSION) -X main.Mode=$(MODE)"

.PHONY: build run build_image

build:
	go build -o ./build/$(TARGET) -v -ldflags=$(LDFLAGS) $(MAIN)

run:
	@echo "Running $(TARGET)"
	@go run -ldflags $(LDFLAGS) $(MAIN)

build_image:
	@echo "Building $(IMAGENAME):$(VERSION)"
	docker build \
	-t $(IMAGENAME):$(VERSION) \
	-t $(IMAGENAME):latest \
	--label "Version=$(VERSION)" \
	--label "Vendor=$(VENDOR)" \
	--label "AppName=$(TARGET)" \
	--no-cache \
	--build-arg VERSION=$(VERSION) \
	--build-arg MAIN=$(MAIN) \
	--build-arg TARGET=$(TARGET) \
	--build-arg MODE=$(MODE) \
	--build-arg LDFLAGS=$(LDFLAGS) \
	--progress=plain \
	-f ./Dockerfile .

run_image:
	docker run -it --rm --name $(TARGET) $(IMAGENAME):$(VERSION)