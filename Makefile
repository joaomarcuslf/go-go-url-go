SHELL := /bin/bash
LINT_VERSION=v1.37.1
GO_BUILD_ENV := CGO_ENABLED=0 GOOS=linux GOARCH=amd64
DOCKER_BUILD=$(shell pwd)/.docker_build
DOCKER_CMD=$(DOCKER_BUILD)/go-getting-started

$(DOCKER_CMD): clean
	mkdir -p $(DOCKER_BUILD)
	$(GO_BUILD_ENV) go build -v -o $(DOCKER_CMD) .

.PHONY: env

env:
	sh scripts/set-local-env.sh

build: env
	. .env && go build main.go

gofmt:
	gofmt -w .

test: env
	docker-compose up -d redis
	sleep 1
	GOLANG_PROTOBUF_REGISTRATION_CONFLICT=ignore
	go test -cover ./...
	docker-compose down

docker-start:
	docker-compose up -d

docker-stop:
	docker-compose down

deploy:
	git tag -a v$(version) -m "v$(version)"
	gcloud app deploy app.yaml -v v$(version)
	git push --tags

clean:
	rm -rf $(DOCKER_BUILD)

deploy: $(DOCKER_CMD)
	git tag -a v$(version) -m "v$(version)"
	git push --tags
	git push heroku main
