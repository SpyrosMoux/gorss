# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.23 AS build-stage

WORKDIR /app

COPY ../go.mod go.sum ./
RUN go mod download

COPY .. .

RUN CGO_ENABLED=0 GOOS=linux go build -o /gorss main.go

# Deploy the application binary into a lean image
FROM alpine:3.21.2 AS build-release-stage
WORKDIR /
COPY --from=build-stage /gorss /gorss
ENTRYPOINT ["/gorss"]
