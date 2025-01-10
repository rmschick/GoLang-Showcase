GOVERSION = $(shell go version | awk '{print $$3;}')
SOURCE_FILES?=./...
GO_DIRECTORIES = $(shell find . -type f -name "*.go" -exec dirname {} \; | sort -u)

export PATH := ./bin:$(PATH)
export CGO_ENABLED := 0

clean:
	rm -rf ./dist ./vendor
.PHONY: clean

upgrade:
	@for dir in $(GO_DIRECTORIES); do \
		echo "Upgrading dependencies in $$dir"; \
		(cd $$dir && go get -t -u ./...); \
	done
.PHONY: upgrade

vendor:
	@for dir in $(GO_DIRECTORIES); do \
		echo "Running go mod vendor in $$dir"; \
		(cd $$dir && go mod vendor); \
	done
.PHONY: vendor

tidy:
	@for dir in $(GO_DIRECTORIES); do \
		echo "Tidying dependencies in $$dir"; \
		(cd $$dir && go mod tidy); \
	done
.PHONY: tidy

lint:
	@for dir in $(GO_DIRECTORIES); do \
		echo "Linting $$dir"; \
		golangci-lint run $$dir --timeout=5m --out-format colored-line-number:stdout || exit 1; \
	done
.PHONY: lint

test:
	@for dir in $(GO_DIRECTORIES); do \
		echo "Testing $$dir"; \
		gotestsum -- -failfast -v -covermode count -timeout 5m $$dir/... || exit 1; \
	done
.PHONY: test

docs:
	# Docs available at http://localhost:6060
	godoc -http=:6060
.PHONY: docs
