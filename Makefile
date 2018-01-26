.PHONY: all
all: build test

#
# Logging
#

### Colour Definitions
END_COLOR=\x1b[0m
GREEN_COLOR=\x1b[32;01m
RED_COLOR=\x1b[31;01m
YELLOW_COLOR=\x1b[33;01m

### End output
end:
	@echo "$(YELLOW_COLOR)🙏  🙏  🙏$(END_COLOR)"

#
# Project Initialisation
#

### Name of the executable, it's possible to have multiple executables
APP_EXECUTABLE="samplecli"

### Get a list of all golang packages
ALL_PACKAGES=$(go list ./... | grep -v "vendor")

#
# Recipes for building existing projects
#

### Clean temporary files
clean:
	@echo "$(GREEN_COLOR)Cleaning unwanted files $(END_COLOR)"
	rm -rf application.toml
	rm -rf coverage.txt
	rm -rf coverage.html
	rm -rf bin/

### First initialisation of a new project on a Mac
init_mac:
	@echo "$(GREEN_COLOR)Initialising dep for Mac $(END_COLOR)"
	brew install dep
	dep init

### First initialisation of a new project on a Linux machine
init_linux:
	@echo "$(GREEN_COLOR)Initialising dep for Linux $(END_COLOR)"
	go get -u github.com/golang/dep/cmd/dep
	dep init

### Build a project for the first time on a Mac
setup_mac:
	@echo "$(GREEN_COLOR)Setting up dep for Mac $(END_COLOR)"
	brew upgrade dep

### Build a project for the first time on a Linux machine
setup_linux:
	@echo "$(GREEN_COLOR)Setting up dep for Linux $(END_COLOR)"
	go get -u github.com/golang/dep/cmd/dep

### Update dependencies
update:
	@echo "$(GREEN_COLOR)Running dep ensure $(END_COLOR)"
	dep ensure

### Fix formatting
fmt:
	@echo "$(GREEN_COLOR)Running fmt $(END_COLOR)"
	go fmt ./...

### Run go vet
vet:
	@echo "$(GREEN_COLOR)Running vet $(END_COLOR)"
	go vet ./...

### Check for linting issues
lint:
	@echo "$(GREEN_COLOR)Running lint $(END_COLOR)"
	@echo "$(RED_COLOR)Linting is not running, fix in Makefile $(END_COLOR)"
	#golint ./... | grep -v vendor

### Copy config from template
copy-config:
	@echo "$(GREEN_COLOR)Copying config from sample $(END_COLOR)"
	cp application.toml.sample application.toml

### Manually test all packages
test:
	@echo "$(GREEN_COLOR)Running tests for all packages $(END_COLOR)"
	go test ./... -v -p=5 -race -covermode=atomic -timeout=30s

### Compile a linux and mac binary in the ./bin folder
compile:
	@echo "$(GREEN_COLOR)Compiling linux and mac binaries in ./bin $(END_COLOR)"
	mkdir -p bin/
	go build -o bin/$(APP_EXECUTABLE) ./cmd/$(APP_EXECUTABLE)
	CGO_ENABLED=0 GOOS=linux go build -o bin/$(APP_EXECUTABLE)_linux ./cmd/$(APP_EXECUTABLE)

### Calculate test coverage for the whole project (except vendors)
coverage:
	@echo "$(GREEN_COLOR)Calculating test coverage across packages $(END_COLOR)"
	echo 'mode: atomic' > coverage.txt && echo '' > coverage.tmp && go list ./... | xargs -n1 -I{} sh -c 'go test -p=5 -race -covermode=atomic -coverprofile=coverage.tmp -timeout=30s {} && tail -n +2 coverage.tmp >> coverage.txt'
	rm coverage.tmp
	go tool cover -html=coverage.txt -o coverage.html
	@echo "$(YELLOW_COLOR)Run open ./coverage.html to view coverage $(END_COLOR)"

### Install all binaries (Repo could have multiple binaries)
install:
	@echo "$(GREEN_COLOR)Installing all binaries $(END_COLOR)"
	go install ./...

### Build the latest source on a mac
build: fmt vet lint coverage install end

### Build the latest source for the first time on a mac
build_fresh: clean setup_mac update fmt vet lint copy-config coverage compile install end

### Build on the CI (usually travisCI)
build_ci: clean setup_linux update fmt vet lint copy-config coverage compile end

#
# Receipes for docker
#

build_docker:
	@echo "$(GREEN_COLOR)Building a docker image $(END_COLOR)"
	docker build -t sudhanshuraheja/sample .

#
# Recipes for starting new projects
#