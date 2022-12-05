migrateup:
	migrate -path db/migration -database "postgres://jnuma:jnuma@localhost:5432/bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://jnuma:jnuma@localhost:5432/bank?sslmode=disable" -verbose down

.PHONY: migrateup migratedown