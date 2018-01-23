.PHONY: all
all: build test

## Project Initialisation

APP_EXECUTABLE="golang-sample"

ALL_PACKAGES=$(go list ./... | grep -v "vendor")

## Recipes for building existing projects

clean:
	rm -rf application.toml
	rm -rf coverage.txt
	rm -rf coverage.html

setup_mac:
	brew install dep
	dep init

setup_linux:
	go get -u github.com/golang/dep/cmd/dep
	dep init

update:
	dep ensure

fmt:
	go fmt ./...

vet:
	go vet ./...

lint:
	#TODO find out why does this always fail
	#golint ./... | grep -v vendor

copy-config:
	cp application.toml.sample application.toml

test:
	go test ./... -v -p=5 -race -covermode=atomic -timeout=30s

compile:
	mkdir -p bin/
	go build -o bin/$(APP_EXECUTABLE)
	CGO_ENABLED=0 GOOS=linux go build -o bin/$(APP_EXECUTABLE)_linux

coverage:
	echo 'mode: atomic' > coverage.txt && echo '' > coverage.tmp && go list ./... | xargs -n1 -I{} sh -c 'go test -covermode=atomic -coverprofile=coverage.tmp {} && tail -n +2 coverage.tmp >> coverage.txt'
	rm coverage.tmp
	go tool cover -html=coverage.txt -o coverage.html

install:
	go install

build: fmt vet lint test coverage install

build_fresh: clean setup_mac update fmt vet lint copy-config test compile coverage install

build_ci: clean setup_linux update fmt vet lint copy-config test compile coverage

## Recipes for starting new projects

#TODO