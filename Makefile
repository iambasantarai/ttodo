BIN_DIR := bin

build:
	go build -o bin/ttodo
run: build
	./bin/ttodo
test: build
	go test -v ./...
clean:
	rm -rf $(BIN_DIR)
