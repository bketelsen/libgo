GO_BIN ?= go

.PHONY: install
install:
	@$(GO_BIN) install -v ./...
	@make tidy

.PHONY: tidy
tidy:
	@$(GO_BIN) mod tidy

.PHONY: deps
deps:
	@$(GO_BIN) get github.com/gobuffalo/release
	@$(GO_BIN) get github.com/gobuffalo/shoulders
	@$(GO_BIN) get -t ./...
	@make tidy

.PHONY: build
build:
	@$(GO_BIN) build -v ./...
	@make tidy

.PHONY: test
test:
	@$(GO_BIN) test ./...
	@make tidy


.PHONY: ci-deps
ci-deps:
	$(GO_BIN) get -tags -t ./...

.PHONY: ci-test
ci-test:
	$(GO_BIN) test -tags -race ./...

.PHONY: lint
lint:
	@gometalinter --vendor ./... --deadline=1m --skip=internal
	@make tidy

.PHONY: update
update:
	@$(GO_BIN) get -u
	@make tidy
	@make test
	@make install
	@make tidy
