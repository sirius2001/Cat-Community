GOOS := $(shell go env GOOS)
GO_VERSION := $(shell go version | awk '{print $$3}')
APP_NAME := app

.PHONY: check-go-version
check-go-version:
    @if [ "$(shell expr $(GO_VERSION) \>= 'go1.20')" = "0" ]; then \
        echo "Go version $(GO_VERSION) does not support Go Modules."; \
        exit 1; \
    fi
	
.PHONY: init
init:
	@go mod tidy

.PHONY: build
build:check-go-version
	@go build -o $(APP_NAME) ./app

.PHONY: run
run: build
	@ sudo  ./app/$(APP_NAME)

.PHONY: clean
clean:
	@rm -f $(APP_NAME)
