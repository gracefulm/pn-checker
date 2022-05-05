# Constants
CMD_NAME=pn-checker
ENTORY_POINT=./main.go
BUILD_DIR=./build/

# Get the packages needed to run the program
.PHONY: setup
setup:
	go mod download
	go mod tidy

# Update dependent packages
.PHONY: update
update:
	go get -u -d ./...

# Build command
.PHONY: build
build: clean-built setup test
	@echo "> building a command..."
	test -e $(BUILD_DIR) || mkdir $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(CMD_NAME) $(ENTORY_POINT)

# Remove files under build directory
.PHONY: clean-built
clean-built:
	@echo "> cleaning binaries under $(BUILD_DIR)..."
	test ! -e $(BUILD_DIR) || rm -rf $(BUILD_DIR)

.PHONY: test
test:
	@echo "> testing source code..."
	go test -cover ./...
