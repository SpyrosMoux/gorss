DOCKER_REPO = ghcr.io/spyrosmoux/gorss
TAG ?= latest
PLATFORM = linux/amd64
BUILD_DIR = build

.PHONY: build build-docker run run-docker

default: run

build:
	go build -o $(BUILD_DIR)/gorss ./main.go

build-docker:
	docker build -t $(DOCKER_REPO)/gorss:$(TAG) . -f docker/Dockerfile --platform $(PLATFORM)

run:
	go run main.go

run-docker: build-docker
	docker run -d $(DOCKER_REPO)/gorss:$(TAG) -p 8080:8080
