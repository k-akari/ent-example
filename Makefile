migrate:
	go run ./cmd/migration/main.go

app:
	go run ./cmd/app/main.go

lint:
	@staticcheck ./... && go vet ./...

.PHONY: migrate app lint