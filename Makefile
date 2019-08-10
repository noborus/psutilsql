BINARY_NAME := psutilsql
SRCS := $(shell git ls-files '*.go')

all: build

test: $(SRCS)
	go test ./...

build: $(BINARY_NAME)

$(BINARY_NAME): $(SRCS)
	go build -o $(BINARY_NAME) ./cmd/psutilsql

install:
	go install ./cmd/psutilsql

clean:
	rm -f $(BINARY_NAME)

.PHONY: all test build install clean
