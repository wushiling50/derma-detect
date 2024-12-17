DIR = $(shell pwd)

.PHONY: env-up
env-up:
	sh init.sh
	@ docker compose -f ./docker/docker-compose.yml up -d

.PHONY: env-down
env-down:
	@ cd ./docker && docker compose down

.PHONY: run
run:
	go build -o derma && ./derma

.PHONY: update
update:
	hz update -idl api.thrift