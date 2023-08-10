BIN_DIR := ./bin
PATH_MAIN_FILE:=./cmd/auth/main.go

build:
	mkdir -p $(BIN_DIR)
	go build $(BUILD_FLAGS) -o $(BIN_DIR)/myapp $(PATH_MAIN_FILE)

test:
	go test -v ./...
runMain:
	go run $(PATH_MAIN_FILE)
build-run:
	mkdir -p $(BIN_DIR)
	go build $(BUILD_FLAGS) -o $(BIN_DIR)/myapp $(PATH_MAIN_FILE)
	$(BIN_DIR)/myapp
.PHONY: build test
