# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

SHELL = /bin/bash
OS = $(shell uname | tr A-Z a-z)

.PHONY: up
up: start ## Spin up the environment

.PHONY: down
down: ## Destroy the environment
	docker compose down -v
	@ if [[ "$$OSTYPE" == "linux-gnu" ]]; then sudo rm -rf var/docker/volumes/; else rm -rf var/docker/volumes/; fi

docker-compose.override.yml: ## Create docker compose override file
	@ if [[ "$$OSTYPE" == "linux-gnu" ]]; then cat docker-compose.override.yml.dist | sed -e 's/# user: "$${uid}:$${gid}"/user: "$(shell id -u):$(shell id -g)"/' > docker-compose.override.yml; else cp docker-compose.override.yml.dist docker-compose.override.yml; fi

.PHONY: start
start: docker-compose.override.yml ## Start services
	@ if [ docker-compose.override.yml -ot docker-compose.override.yml.dist ]; then diff -u docker-compose.override.yml* || (echo "!!! The distributed docker-compose.override.yml example changed. Please update your file accordingly (or at least touch it). !!!" && false); fi
	mkdir -p var/docker/volumes/postgres
	docker compose up -d

.PHONY: stop
stop: ## Stop services
	docker compose stop

.PHONY: ps
ps: ## Check the status of services services
	docker compose ps

.PHONY: shell
shell: ## Start a shell with the Temporal CLI
	docker compose exec --profile cli temporal-admin-tools bash

.PHONY: worker
worker: ## Start the worker
	go run ./cmd/worker/

.PHONY: test
test: ## Run tests
	go test ./...

.PHONY: slides
slides: ## Open slides in the browser
	reveal-md --template etc/reveal/template/reveal.html --theme slides/theme.css --css slides/custom.css slides/index.md -w

.PHONY: build-slides
build-slides: ABSOLUTE_URL := http://localhost
build-slides:
	reveal-md --template etc/reveal/template/reveal.html --theme slides/theme.css --css slides/custom.css --static public/ --static-dirs slides/assets --absolute-url ${ABSOLUTE_URL} slides/index.md

.PHONY: print-slides
print-slides:
	decktape --size='2048x1536' https://sagikazarmark.github.io/temporal-intro-workshop slides.pdf

.PHONY: deps
deps:
	npm install -g reveal-md

.PHONY: help
.DEFAULT_GOAL := help
help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-10s\033[0m %s\n", $$1, $$2}'
