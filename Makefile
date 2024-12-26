build:
	@go build -o bin/blogo ./cmd/blogo/

run: build
	@./bin/blogo -cfg ./config/config.json

tidy:
	@go mod tidy

test:
	@go test ./... -v