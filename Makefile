build:
	@cd web/fibonacci && npm run build

test:
	@go test ./... -v
