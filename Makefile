build:
	go build -o bin/ttodo
run: build
	./bin/ttodo
test: build
	go test -v ./...
