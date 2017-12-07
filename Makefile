NAME := commango
DEVEL := development
PROD := production

dev: dep install config-$(DEVEL) build
pro: dep install config-$(DEVEL) build

dep:
	dep ensure

install:
	go install

build:
	GOOS=linux GOARCH=amd64 go build -tags netgo -installsuffix netgo -o bin/$(NAME)

config-%:
	cp config.$*.toml bin/config.toml

run-dev:
	go run main.go

build-image-dev: dev
	sh docker-build.sh

build-image-pro: pro
	sh docker-build.sh

docker-push:
	sh docker-push.sh

docker-run:
	sh docker-run.sh