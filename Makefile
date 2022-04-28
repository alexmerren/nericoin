GO := go

NAME := nericoin
CMD_DIR := $(CURDIR)/cmd
BIN_DIR := $(CURDIR)/bin
MAIN_LOCATION := $(CMD_DIR)/$(NAME)/main.go


.PHONY: build
build:
	$(GO) build -o $(BIN_DIR)/$(NAME) -mod=vendor $(MAIN_LOCATION)

.PHONY: run 
run:
	$(BIN_DIR)/$(NAME)

.PHONY: vendor 
vendor:
	$(GO) mod tidy
	$(GO) mod vendor

.PHONY: install-tools 
install-tools:
	$(GO) get github.com/boltdb/bolt/...
