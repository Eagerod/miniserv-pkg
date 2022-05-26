GO := go

MAIN_FILE := main.go

BUILD_DIR := build
EXECUTABLE := serv
BIN_NAME := $(BUILD_DIR)/$(EXECUTABLE)

COVERAGE_FILE=./coverage.out

ALL_GO_DIRS = $(shell find . -iname "*.go" -exec dirname {} \; | sort | uniq)
SRC := $(shell find . -iname "*.go" -and -not -name "*_test.go")
SRC_WITH_TESTS := $(shell find . -iname "*.go")


.PHONY: all
all: $(BIN_NAME)

$(BIN_NAME): $(SRC)
	@mkdir -p $(BUILD_DIR)
	$(GO) build -o $(BIN_NAME)


.PHONY: test
test: $(SRC)
	@if [ -z $$T ]; then \
		$(GO) test -v ./...; \
	else \
		$(GO) test -v ./... -run $$T; \
	fi


$(COVERAGE_FILE): $(SRC_WITH_TESTS)
	$(GO) test -v --coverprofile=$(COVERAGE_FILE) ./...

.PHONY: coverage
coverage: $(COVERAGE_FILE)
	$(GO) tool cover -func=$(COVERAGE_FILE)

.PHONY: pretty-coverage
pretty-coverage: $(COVERAGE_FILE)
	$(GO) tool cover -html=$(COVERAGE_FILE)


# Project structure prevents go fmt ./... from working?
.PHONY: fmt
fmt:
	@find . -type f -iname "*.go" -exec go fmt {} \;


.PHONY: clean
clean:
	rm -rf $(COVERAGE_FILE) $(BUILD_DIR)
