DC := @docker-compose

.PHONY: up
up:
	$(DC) up

.PHONY: upd
upd:
	$(DC) up -d

.PHONY: build
build:
	$(DC) build

.PHONY: down
down:
	$(DC) down

.PHONY: api
api:
	$(DC) exec api bash

.PHONY: doc
doc:
	$(DC) exec api godoc -http=:9000

.PHONY: front
front:
	$(DC) exec front sh

.PHONY: db
db:
	$(DC) exec db bash -c 'mysql -u $$MYSQL_USER -p$$MYSQL_PASSWORD $$MYSQL_DATABASE'

.PHONY: seed
seed:
	$(DC) exec api go run infrastructure/seed/seed.go

.PHONY: init
init:
	$(DC) up -d --build
	$(DC) exec -T front sh -c 'cp .env.local .env'
	$(DC) exec -T api bash -c 'go mod tidy'