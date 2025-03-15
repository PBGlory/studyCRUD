DB_DSN := "postgres://postgres:yourpassword@localhost:5432/main?sslmode=disable"
MIGRATE := migrate - path.migration - database $(DB_DSN)

migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}


migrate:
	$(MIGRATE) up


migrate-down:
	$(MIGRATE) down


run:
	go run cmd/app/ main.go