all: build

help:                                  ## Display this help message
	@echo "Please use \`make <target>\` where <target> is one of:"
	@grep '^[a-zA-Z]' $(MAKEFILE_LIST) | \
		awk -F ':.*?## ' 'NF==2 {printf "  %-26s%s\n", $$1, $$2}'

build-local:                           ## Build `aleksi/test_db:local`
	docker buildx build --pull --tag=aleksi/test_db:local --load .

build-push:                            ## Build aleksi/test_db:$(DOCKER_TAG) and push
	test $(DOCKER_TAG)
	docker buildx build --pull --platform=linux/arm64/v8,linux/amd64 --tag=aleksi/test_db:$(DOCKER_TAG) --push .

up:                                    ## Start environment
	docker compose up --always-recreate-deps --force-recreate --remove-orphans --renew-anon-volumes --build

mysql:                                 ## Start mysql shell
	docker compose exec mysql mysql --database sakila

psql:                                  ## Start psql shell
	docker compose exec postgresql psql --username postgres --dbname pagila

mongosh:                               ## Start mongosh shell
	docker compose exec mongodb mongosh mongodb://localhost/monila

.PHONY: mysql
