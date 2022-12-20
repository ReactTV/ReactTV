SHELL := /bin/bash

default: help

help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

version: ## Version prints out the next semantic version
	@npx -p @semantic-release/exec -p semantic-release semantic-release --dry-run

login: ## Authenticates with Docker Registry (example: make login username=$DOCKER_USER password=$DOCKER_PASSWORD)
	@echo "Authenticating with docker registry (ghcr.io)..." ; docker login ghcr.io --username $(username) --password $(password)

build: ## Builds all services with a tagged version (example: make build version=v1.0.0 commit=`git rev-parse HEAD` )
	@cd deploy/scripts ; source docker.sh ; build $(version) $(commit)

push: login ## Push all services with a tagged version (example: make build version=v1.0.0 username=$DOCKER_USER password=$DOCKER_PASSWORD)
	@cd deploy/scripts ; source docker.sh ; push $(version)

dev: ## Run ReactTV in Development Mode
	@tilt up

api-docs: ## Generate API documentation for ReactTV/Server
	@cd services/server/pkg/api && swag init -g api.go --parseDependency true