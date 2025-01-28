BIN_DIR := bin

build:
	go build -o $(BIN_DIR)/ttodo
run: build
	./$(BIN_DIR)/ttodo
test: build
	go test -v ./...
clean:
	rm -rf $(BIN_DIR)
