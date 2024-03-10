mock:
	mockery --all --keeptree
migrate:
	migrate -source file://database/migrations \
			-database postgres://admin:admin@127.0.0.1:5432/twitter?sslmode=disable up

rollback:
	migrate -source file://database/migrations \
			-database postgres://admin:admin@127.0.0.1:5432/twitter?sslmode=disable down 1

drop:
	migrate -source file://database/migrations \
			-database postgres://admin:admin@127.0.0.1:5432/twitter?sslmode=disable drop

migration:
	@read -p "Enter migration name: " name; \
		migrate create -ext sql -dir database/migrations $$name

run:
	go run cmd/graphqlserver/*.go

generate:
	go generate ./..