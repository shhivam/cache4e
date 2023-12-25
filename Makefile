migrate:
	migrate -path "./internal/app/postgres/migrations" -database "postgres://postgres:india123@0.0.0.0:5432/acme?sslmode=disable" up

dbs-up:
	docker compose up -d

dbs-down:
	docker compose down

run:
	make dbs-up
	make migrate
	go run cmd/app/app.go

run-server:
	go run cmd/app/app.go


