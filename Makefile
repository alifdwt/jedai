postgres:
	docker run --name postgres16 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root jedai_db

dropdb:
	docker exec -it postgres16 dropdb jedai_db

migrateup:
	migrate -path server/db/migration -database "postgresql://root:secret@localhost:5432/jedai_db?sslmode=disable" -verbose up

migratedown:
	migrate -path server/db/migration -database "postgresql://root:secret@localhost:5432/jedai_db?sslmode=disable" -verbose down

sqlc:
	cd server && sqlc generate

test:
	cd server && go test -v -cover -short ./...

backend:
	cd server && go run main.go

frontend:
	cd client && npm run dev

.PHONY: postgres createdb dropdb migrateup migratedown