postgres:
	docker run --name postgres-ecommerce -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:latest

createdb:
	docker exec -it postgres-ecommerce createdb --username=root --owner=root ecommerce-db

dropdb:
	docker exec -it postgres-ecommerce dropdb ecommerce-db


.PHONY: postgres createdb dropdb



migrateup:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/ecommerce-db?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/ecommerce-db?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/ecommerce-db?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/ecommerce-db?sslmode=disable" -verbose down 1

migrateversion:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/ecommerce-db?sslmode=disable" -verbose version

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go bank-backend-project/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server mock