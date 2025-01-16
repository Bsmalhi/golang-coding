build:
	go build -o bin/$(shell basename $(PWD)) cmd/golangCoding/main.go
run:
	go run cmd/golangCoding/main.go
test:
	go test -v ./...