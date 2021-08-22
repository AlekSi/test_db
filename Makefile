build:
	docker build --pull --squash --tag aleksi/test_db:local .

mysql:
	docker compose exec mysql mysql --database sakila

psql:
	docker compose exec postgresql psql --username postgres --dbname pagila
