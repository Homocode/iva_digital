dockernetwork:
	docker network create iva_digital_network

postgres:
	docker run -p 5432:5432 --name postgres --network iva_digital_network -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123 -d postgres

createdb:
	docker exec -it postgres createdb --username=root --owner=root contable

dropdb:
	docker exec -it postgres dropdb --username=root contable

migrateup:
	migrate -path db/migration -database "postgresql://root:123@localhost:5432/contable?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:123@localhost:5432/contable?sslmode=disable" -verbose down

server:
	docker run --name iva_digital_api --network iva_digital_network -p 8080:8080 iva_digital_api:latest
	
.PHONY: postgres createdb dropdb migrateup migratedown