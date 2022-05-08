migrate:
	go run ./cmd/migration/main.go

app:
	go run ./cmd/app/main.go

.PHONY: migrate app