postgres:
	docker run --name postgres15 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

create_db:
	docker exec -it postgres15 createdb --username=root --owner=root simple_bank

drop_db:
	docker exec -it postgres15 dropdb simple_bank

migrate_up:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable" -verbose up

migrate_down:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable" -verbose down

sqlc_init:
	docker run --rm -v $(CURDIR):/src -w /src kjconroy/sqlc init

sqlc_generate:
	docker run --rm -v $(CURDIR):/src -w /src kjconroy/sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY:
	postgres
	create_db
	drop_db
	migrate_up
	migrate_down
	sqlc_init
	sqlc_generate
	test
	server