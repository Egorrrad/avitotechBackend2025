lint:
	@golangci-lint run ./...

generate:
	oapi-codegen --config=api/oapi-codegen.yaml api/swagger.yaml