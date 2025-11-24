DOCKER_REPO = ghcr.io/SpyrosMoux/gorss
TAG ?= latest
PLATFORM = linux/amd64
BUILD_DIR = build

.PHONY: build build-docker-api build-docker-web run run-docker

default: run

build:
	go build -o $(BUILD_DIR)/gorss ./main.go

build-docker-api:
	docker build -t $(DOCKER_REPO)/gorss:$(TAG) . -f docker/Dockerfile.api --platform $(PLATFORM)
	docker push $(DOCKER_REPO)/gorss:$(TAG)

build-docker-web:
	docker build -t $(DOCKER_REPO)/gorss-web:$(TAG) . -f docker/Dockerfile.web --platform $(PLATFORM)
	docker push $(DOCKER_REPO)/gorss-web:$(TAG)

run:
	go run main.go
