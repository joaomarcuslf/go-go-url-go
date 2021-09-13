SHELL := /bin/bash
LINT_VERSION=v1.37.1

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
