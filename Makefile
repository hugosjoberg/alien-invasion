lint:
	golangci-lint run

test:
	go test -v ./...

.PHONY : example
example:
	go run main.go -n 3 -p "example/map1.txt"
