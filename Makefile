.PHONY: build test exec

PROJECT_NAME=$(shell cat .env | awk -F= 'NR==1{print $$2}')

build:
	echo "PROJECT_NAME=$(PROJECT_NAME)" > .env
	docker-compose up -d --build --force-recreate

test:
	docker-compose run --rm --name go-test go go $(PROJECT_NAME) -v ./...

exec:
	docker exec -it $(PROJECT_NAME) bash
