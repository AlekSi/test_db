all: build

build-local:
	docker buildx build --pull --tag=aleksi/test_db:local --load .

build-push:
	test $(DOCKER_TAG)
	docker buildx build --pull --platform=linux/arm64/v8,linux/amd64 --tag=aleksi/test_db:$(DOCKER_TAG) --push .

up:
	docker compose up --always-recreate-deps --force-recreate --remove-orphans --renew-anon-volumes --build

mysql:
	docker compose exec mysql mysql --database sakila

psql:
	docker compose exec postgresql psql --username postgres --dbname pagila

mongosh:
	docker compose exec mongodb mongosh mongodb://localhost/monila

.PHONY: mysql
