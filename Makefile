BINARY_NAME := psutilsql
SRCS := $(shell git ls-files '*.go')
LDFLAGS := "-X main.version=$(shell git describe --tags --abbrev=0 --always) -X main.revision=$(shell git rev-parse --short HEAD)"

all: build

test: $(SRCS)
	go test ./...

build: $(BINARY_NAME)

$(BINARY_NAME): $(SRCS)
	CGO_ENABLED=0 go build -ldflags $(LDFLAGS) -o $(BINARY_NAME) ./cmd/psutilsql

install:
	CGO_ENABLED=0 go install -ldflags $(LDFLAGS) ./cmd/psutilsql

clean:
	rm -f $(BINARY_NAME)

.PHONY: all test build install clean
