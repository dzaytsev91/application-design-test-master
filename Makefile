export GO111MODULE=on

LOCAL_BIN := $(CURDIR)/bin
GOLANGCI_TAG ?= 1.54.2
GOLANGCI_BIN := $(LOCAL_BIN)/golangci-lint

install-lint:
ifeq ($(wildcard $(GOLANGCI_BIN)),)
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v$(GOLANGCI_TAG)
# Устанавливаем текущий путь для исполняемого файла линтера.
GOLANGCI_BIN := $(LOCAL_BIN)/golangci-lint
else
	$(info Golangci-lint is already installed to $(GOLANGCI_BIN))
endif

lint: install-lint
	$(info Running lint against changed files...)
	$(GOLANGCI_BIN) run \
		--new-from-rev=origin/master \
		--config=.golangci.pipeline.yaml \
		./...

lint-full: install-lint
	$(GOLANGCI_BIN) run \
		--config=.golangci.pipeline.yaml \
		./...


test:
	$(info Running tests...)
	go test ./...


build:
	go build \
		-cover \
		-covermode=atomic \
		-ldflags "$(LDFLAGS)" \
		-o "$(LOCAL_BIN)/book-api" \
		$(PGO_FLAG) \
		./cmd/book-api

lint-fix-affected-files: install-lint
	$(info Trying to autofix linting for files affected by changes since origin/master ...)
	$(GOLANGCI_BIN) run \
		--new-from-rev=origin/master \
		--whole-files \
		--fix \
		--config=.golangci.pipeline.yaml \
		./...