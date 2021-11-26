all: build

build-local:
	docker buildx build --pull --tag=aleksi/test_db:local --load .

build-push:
	docker buildx build --pull --platform=linux/arm64/v8,linux/amd64 --tag=aleksi/test_db:latest --push .

up:
	docker compose up --always-recreate-deps --force-recreate --remove-orphans --renew-anon-volumes --build

mysql:
	docker compose exec mysql mysql --database sakila

psql:
	docker compose exec postgresql psql --username postgres --dbname pagila

mongosh:
	docker compose exec mongodb mongosh mongodb://localhost/monila

.PHONY: mysql
