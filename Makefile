SOURCE_FILES?=./...
TEST_PATTERN?=.
TEST_OPTIONS?=

export PATH := ./bin:$(PATH)
export GO111MODULE := on

# Install all the build and lint dependencies
setup:
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh
	go mod download

# Run all the tests
test:
	go test $(TEST_OPTIONS) -covermode=atomic -coverprofile=coverage.txt ./... -run $(TEST_PATTERN) -timeout=30s

# Run all the tests and opens the coverage report
cover: test
	go tool cover -html=coverage.txt

# gofmt and goimports all go files
fmt:
	find . -name '*.go' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done

# Run all the linters
lint:
	golangci-lint run --enable-all ./...

# Run shcheck against itself
run:
	go run main.go --ignore='sh/testdata/*.sh'

# Run all the tests and code checks
ci: run lint test

# Build a local version
build:
	go build .

# Install to $GOPATH/src
install:
	go install

.DEFAULT_GOAL := build
