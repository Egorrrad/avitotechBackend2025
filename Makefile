lint:
	@golangci-lint run ./...

generate:
	oapi-codegen --config=api/oapi-codegen.yaml api/swagger.yaml

run-service:
	docker compose up --build -d
stop-service:
	docker-compose stop
rm-service:
	docker-compose down
run-tests:
	go test -v -cover ./...
test-cov:
	go test -coverprofile=coverage.out ./... & go tool cover -html=coverage.out