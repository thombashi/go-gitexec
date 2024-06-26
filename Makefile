VERSION := v0.0.2

BIN_DIR := $(CURDIR)/bin

STATICCHECK := $(BIN_DIR)/staticcheck
TESTIFYILINT := $(BIN_DIR)/testifylint

$(BIN_DIR):
	mkdir -p $(BIN_DIR)

$(STATICCHECK): $(BIN_DIR)
	GOBIN=$(BIN_DIR) go install honnef.co/go/tools/cmd/staticcheck@latest

$(TESTIFYILINT): $(BIN_DIR)
	GOBIN=$(BIN_DIR) go install github.com/Antonboom/testifylint@latest

.PHONY: clean
clean:
	rm -rf $(BIN_DIR)

.PHONY: check
check: $(STATICCHECK) $(TESTIFYILINT)
	$(STATICCHECK) ./...
	$(TESTIFYILINT) ./...

.PHONEY: fmt
fmt: $(TESTIFYILINT)
	gofmt -w -s .
	$(TESTIFYILINT) -fix ./...

.PHONEY: push-tag
push-tag:
	git push origin $(VERSION)

.PHONY: tag
tag:
	git tag $(VERSION) -m "Release $(VERSION): Signed with gitsign"

.PHONY: test
test:
	go test -v ./...
