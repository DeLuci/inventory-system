
network:
	docker network create inventory-network

postgres:
	docker run --name postgres14 --network inventory-network -p 5432:5432 -e POSTGRES_PASSWORD=dummypassword -e POSTGRES_USER=root -d postgres:14-alpine

createdb:
	docker exec -it postgres14 createdb --username=root --owner=root inventory

dropdb:
	docker exec -it postgres14 dropdb inventory

migrateup:
	migrate -path db/migration -database "postgresql://root:dummypassword@localhost:5432/inventory?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:dummypassword@localhost:5432/inventory?sslmode=disable" -verbose down
.PHONY: postgres createdb dropdb migrateup migrationdown network