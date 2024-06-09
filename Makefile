build:
	@go build -o bin/async-jobs

run: build
	@./bin/async-jobs

watch: 
	@go run main.go

test:
	@go test ./... -v
