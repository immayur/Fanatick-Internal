# tests the repository packages
test:
	-dropdb $(DB_NAME)
	createdb $(DB_NAME)
	migrate -path=./fanatick/postgres/migrations -database="postgres://localhost/$(DB_NAME)?sslmode=disable" up
	DB_URL=postgres://localhost/$(DB_NAME)?sslmode=disable go test ./...

# makes a migration file
migration:
	migrate create -ext "sql" -dir "./fanatick/postgres/migrations" $(name)

# runs the migrations
migrate:
	migrate -path=./fanatick/postgres/migrations -database=$(DB_URL) up
