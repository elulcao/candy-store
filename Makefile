BIN_NAME=candy_store
UNAME_S := $(shell uname -s)

ifeq ($(UNAME_S),Linux)
	goos_type = linux
endif

ifeq ($(UNAME_S),Darwin)
	goos_type = darwin
endif

all: test vet fmt clean build

test:
	go test ./...

vet:
	go vet ./...

fmt:
	go fmt ./...

clean: tidy
	go clean
	rm -f ${BIN_NAME}.$(goos_type)

tidy:
	go mod tidy
	go mod vendor

build: tidy
	env GO111MODULE=on GOOS=$(goos_type) GOARCH=amd64 go build -v -o "$(BIN_NAME).$(goos_type)"

# Release for linux and darwin
release:
	mkdir -p releases
	env GO111MODULE=on GOOS=linux  GOARCH=amd64 go build -v -o releases/$(BIN_NAME).linux
	env GO111MODULE=on GOOS=darwin GOARCH=amd64 go build -v -o releases/$(BIN_NAME).darwin

install: build
	install -m 0755 "$(BIN_NAME).$(goos_type)" "/usr/local/bin/$(BIN_NAME)"