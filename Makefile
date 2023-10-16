IMAGE_NAME_SERVICE = mfv-codingchallenge
VERSION_SERVICE = 1.0

HOST = 127.0.0.1
PORT = 3306
DATABASE = mfv
USER = root
PASSWORD = secret

init-db:
	docker exec -it database mysql -u root -psecret -e "CREATE DATABASE mfv;"
.PHONY: init-db

migrate-up:
	migrate -source "file://migrations" -database "mysql://$(USER):$(PASSWORD)@tcp($(HOST):$(PORT))/$(DATABASE)" up
.PHONY: migration-up

build:
	go build -o bin/runtime main.go
.PHONY: build

build-docker:
	docker build -t $(IMAGE_NAME_SERVICE):$(VERSION_SERVICE) --force-rm -f Dockerfile .
.PHONY: build-docker

start:
	docker run -it --name $(IMAGE_NAME_SERVICE) -p 9090:9090 $(IMAGE_NAME_SERVICE):$(VERSION_SERVICE)
.PHONY: start

start-docker-compose:
	docker-compose up -d
.PHONY: start-docker-compose

mock:
	cd internal && rm -r mocks && mockery --all --keeptree --case underscore && cd ..
.PHONY: mock