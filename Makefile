SHELL := /bin/bash
LINT_VERSION=v1.37.1

gofmt:
	gofmt -w .

test:
	GOLANG_PROTOBUF_REGISTRATION_CONFLICT=ignore go test -cover ./...

docker-start:
	docker-compose up -d

docker-stop:
	docker-compose down
