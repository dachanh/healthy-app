.PHONY: install migrate
# Define all targets.
TARGETS := postgres


install:
	go mod download
	go install -tags "postgres" github.com/golang-migrate/migrate/v4/cmd/migrate

migrateup:
	migrate -path /packages/migrate/postgres/ddl -database 'postgres://root:secret@postgres.arent.healthyapp:5432/develop-healthy-arent?sslmode=disable' -verbose up

migratedown:
	migrate -path /packages/migrate/postgres/ddl -database 'postgres://root:secret@postgres.arent.healthyapp:5432/develop-healthy-arent?sslmode=disable' -verbose down
