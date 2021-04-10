build:
	@go mod download
	@cd web/fibonacci && npm i && npm run build

test:
	@go test ./... -v

run:
	@go run .
