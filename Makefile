postgresinit:
	docker run --name mydb -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:latest

postgres:
	docker exec -it mydb psql

createdb:
	docker exec -it mydb createdb --username=root --owner=root go-chat

dropdb:
	docker exec -t mydb dropdb go-chat

migrateup:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5433/go-chat?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5433/go-chat?sslmode=disable" -verbose down

.PHONY: postgresinit postgres createdb dropdb migrateup migratedown