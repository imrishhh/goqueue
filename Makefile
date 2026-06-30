migration:
	go run ./executor/migration/migration.go $(action)

migration-create:
	goose -dir ./migrations create $(name) sql

run-coordinator:
	go run ./cmd/coordinator/main.go

run-worker:
	go run ./cmd/worker/main.go

up:
	docker compose -f ./deploy/docker-compose.yml up -d
