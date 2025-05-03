DOCKER_COMPOSE = docker compose
ENV ?= dev
DOCKER_COMPOSE_FILE = $(if $(filter prod,$(ENV)),-f docker-compose.prod.yml,)
DOCKER_COMPOSE_CMD = $(DOCKER_COMPOSE) $(DOCKER_COMPOSE_FILE)

.PHONY: up build down stop inapi inweb indb logapi logweb logdb ps reup help

up:
	$(DOCKER_COMPOSE_CMD) up -d

build:
	$(DOCKER_COMPOSE_CMD) build --no-cache

down:
	$(DOCKER_COMPOSE_CMD) down

stop:
	$(DOCKER_COMPOSE_CMD) stop

inapi:
	$(DOCKER_COMPOSE_CMD) exec api bash

inweb:
	$(DOCKER_COMPOSE_CMD) exec web bash

indb:
	$(DOCKER_COMPOSE_CMD) exec db bash

logapi:
	$(DOCKER_COMPOSE_CMD) logs -f api

logweb:
	$(DOCKER_COMPOSE_CMD) logs -f web

logdb:
	$(DOCKER_COMPOSE_CMD) logs -f db

ps:
	$(DOCKER_COMPOSE_CMD) ps

reup: down up

help:
	@echo "Usage: make [target] [ENV=dev|prod]"
	@echo ""
	@echo "Targets:"
	@echo "  up        Start containers in the specified environment (default: dev)"
	@echo "  build     Build containers without cache"
	@echo "  down      Stop and remove containers, networks, and volumes"
	@echo "  stop      Stop containers"
	@echo "  inapi     Access api container via bash"
	@echo "  inweb     Access web container via bash"
	@echo "  indb      Access db container via bash"
	@echo "  logapi    Show logs for the api container"
	@echo "  logweb    Show logs for the web container"
	@echo "  logdb     Show logs for the db container"
	@echo "  ps        Show status for containers"
	@echo "  reup      Re-up containers"