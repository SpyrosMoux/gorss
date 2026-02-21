DOCKER_REPO = ghcr.io/spyrosmoux/gorss
TAG ?= latest
PLATFORM = linux/amd64
BUILD_DIR = build
DB_HOST ?= localhost
DB_PORT ?= 5432
DB_USER ?= gorss
DB_PASS ?= gorss
DB_NAME ?= gorss

default: run

.PHONY: build
build:
	go build -o $(BUILD_DIR)/gorss ./main.go

.PHONY: build-docker-api
build-docker-api:
	docker build -t $(DOCKER_REPO)/gorss:$(TAG) . -f docker/Dockerfile.api --platform $(PLATFORM)
	docker push $(DOCKER_REPO)/gorss:$(TAG)

.PHONY: build-docker-web
build-docker-web:
	docker build -t $(DOCKER_REPO)/gorss-web:$(TAG) . -f docker/Dockerfile.web --platform $(PLATFORM)
	docker push $(DOCKER_REPO)/gorss-web:$(TAG)

.PHONY: dev
dev: atlas-migrate atlas-apply
	docker compose -f docker/docker-compose.yaml up -d
	make run

.PHONY: run
run: atlas-migrate atlas-apply
	go run main.go

.PHONY: atlas-inspect
atlas-inspect:
	atlas schema inspect --env gorm --url "env://src"

.PHONY: atlas-migrate
atlas-migrate:
	atlas migrate diff --env gorm

.PHONY: atlas-apply
atlas-apply:
	atlas schema apply --env gorm \
		-u "postgres://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" \
		$(if $(CI),--auto-approve,)
