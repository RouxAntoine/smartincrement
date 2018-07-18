.ONESHELL:
.PHONY: run

OPT_VERBOSE=-v

# Basic go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOINSTALL=$(GOCMD) install $(OPT_VERBOSE)
GOGET=$(GOCMD) get $(OPT_VERBOSE)

# TARGET := $(shell echo $${PWD\#\#*/})

# go source files, ignore vendor directory
SRC=$(shell find . -type f -name '*.go' -not -path "./vendor/*")

run: build
	@chmod +x ${GOPATH}/bin/$(PROGRAM_NAME)
	@${GOPATH}/bin/$(PROGRAM_NAME) ${PROGRAM_ARG}

build: $(SRC)
	@cd ${PROGRAM_NAME}
	@$(GOGET) -d ./...
	@$(GOINSTALL) ./...

clean:
	$(GOCLEAN)
	rm -f ${GOPATH}/bin/$(PROGRAM_NAME)
