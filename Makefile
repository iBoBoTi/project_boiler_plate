# create migration file
#	migrate create -ext sql -dir db/migrations -seq create_roles_table

include .env
export

run:
	go run cmd/main.go

migrate-up:
	migrate -path db/migrations -database $(DATABASE_URL) -verbose up

migrate-down:
	migrate -path db/migrations -database $(DATABASE_URL) -verbose down

migrate-clean:
	migrate -path db/migrations -database $(DATABASE_URL) force 1