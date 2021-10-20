all: build

build:
	docker build --pull --squash --tag aleksi/test_db:local .

up:
	docker compose up --always-recreate-deps --force-recreate --remove-orphans --renew-anon-volumes --build

mysql:
	docker compose exec mysql mysql --database sakila

psql:
	docker compose exec postgresql psql --username postgres --dbname pagila

mongosh:
	docker compose exec mongodb mongosh mongodb://localhost/monila

.PHONY: mysql
