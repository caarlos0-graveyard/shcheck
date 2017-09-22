TEST_PATTERN?=.
TEST_OPTIONS?=-race

# Install all the build and lint dependencies
setup:
	go get -u github.com/alecthomas/gometalinter
	go get -u github.com/golang/dep/...
	go get -u github.com/pierrre/gotestcover
	go get -u golang.org/x/tools/cmd/cover
	dep ensure
	gometalinter --install --update

# Run all the tests
test:
	gotestcover $(TEST_OPTIONS) -covermode=atomic -coverprofile=coverage.txt ./... -run $(TEST_PATTERN) -timeout=30s

# Run all the tests and opens the coverage report
cover: test
	go tool cover -html=coverage.txt

# gofmt and goimports all go files
fmt:
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done

# Run all the linters
lint:
	gometalinter --vendor ./...

# Run shcheck against itself
run:
	go run main.go --ignore='vendor/**/*' --ignore='sh/testdata/*.sh'

# Run all the tests and code checks
ci: run lint test

# Build a local version
build:
	go build .

# Install to $GOPATH/src
install:
	go install

.DEFAULT_GOAL := build
