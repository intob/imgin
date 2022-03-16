OK_COLOR=\033[32;01m
NO_COLOR=\033[0m

build:
	@echo "$(OK_COLOR)==> Compiling binary$(NO_COLOR)"
	go test && go build -o bin/imgin

test:
	go test

install:
	go get -u .

docker-build:
	@echo "$(OK_COLOR)==> Building Docker image$(NO_COLOR)"
	docker build . --no-cache=true -t intob/imgin:latest

docker-push:
	@echo "$(OK_COLOR)==> Pushing Docker image $(NO_COLOR)"
	docker push intob/imgin:latest

docker: docker-build docker-push

.PHONY: test docker-build docker-push docker