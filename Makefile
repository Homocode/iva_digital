createdb:
	psql postgres -U mel -c 'CREATE DATABASE contable;'

dropdb:
	psql postgres -U mel -c 'DROP DATABASE contable;'

migrateup:
	migrate -path db/migration -database "postgresql://mel@localhost:5432/contable?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://mel@localhost:5432/contable?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown