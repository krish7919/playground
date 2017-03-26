GOCMD=go
GOVET=$(GOCMD) tool vet
GOTEST=$(GOCMD) test
GOCOVER=$(GOCMD) tool cover
#GOINSTALL=$(GOCMD) install -x
GOFMT=gofmt -s -w

PWD=$(shell pwd)
BINARY_PATH=$(PWD)
BINARY_NAME=fsm
MAIN_FILE = $(BINARY_PATH)/fsm.go
SRC_FILES = $(BINARY_PATH)/fsm.go \
			$(BINARY_PATH)/fsm_test.go
TEST_DIR = $(shell dirname $(MAIN_FILE))


.PHONY: all

all: format vet test coverage

format:
	$(GOFMT) $(SRC_FILES)

#clean:
#	@echo "removing any pre-built binary";
#	-@rm $(BINARY_PATH)/$(BINARY_NAME);

#build: format
#	$(shell cd $(BINARY_PATH) && \
#		export GOPATH="$(BINARY_PATH)" && \
#		export GOBIN="$(BINARY_PATH)" && \
#		CGO_ENABLED=0 GOOS=linux $(GOINSTALL) -ldflags "-s" -a \
#		-installsuffix cgo $(MAIN_FILE))

vet:
	$(GOVET) .

test:
	@cd $(TEST_DIR) && export GOPATH="$(BINARY_PATH)" && $(GOTEST) -v

coverage:
	@cd $(TEST_DIR) && export GOPATH="$(BINARY_PATH)" && \
		$(GOTEST) -v -coverprofile=coverage.out && \
		$(GOCOVER) -func=coverage.out && \
		rm coverage.out
