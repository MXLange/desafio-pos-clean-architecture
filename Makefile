COMPOSE := docker compose

.PHONY: up down build rebuild logs ps

up:
	$(COMPOSE) build --no-cache
	$(COMPOSE) up -d

down:
	$(COMPOSE) down -v

build:
	$(COMPOSE) build

rebuild:
	$(COMPOSE) build --no-cache

logs:
	$(COMPOSE) logs -f

ps:
	$(COMPOSE) ps
