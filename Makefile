PROJECT_NAME := "backend-event-tracker-lib"
PKG := "github.com/yofio-mx/$(PROJECT_NAME)"

.PHONY: linter

linter: ## Executes the linter for reviewing the code health
	@echo "Executing linter for golang"
	@golangci-lint run
	@echo "Linter successfully passed"

test: unit-test linter## Execute all tests
	@echo "Test run successfully"

unit-test: ## Run unit tests
	@go test -short ${PKG_LIST}

upgrade-deps: ## Upgrade all dependencies
	@go get -u -t ./...
	@go mod tidy

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36mmake %-15s\033[0m %s\n", $$1, $$2}'
