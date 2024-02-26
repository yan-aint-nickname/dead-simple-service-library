GO ?= go
GOFMT ?= gofmt "-s"
GOFILES := $(shell find . -name "*.go")
GOMODULES := $(shell go list -f '{{.Dir}}' -m | xargs)

.PHONY: test
test:
	$(GO) test

.PHONY: fmt
fmt:
	$(GOFMT) -w $(GOFILES)

lint:
	@hash golangci-lint > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		echo "install golangci-lint"; \
	fi
	for MODULE in $(GOMODULES); do golangci-lint run $$MODULE || exit 1; done;
