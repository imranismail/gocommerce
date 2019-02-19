
serve:
	@go run cmd/serve/main.go
gen.migration:
	@migrate create -dir migrations -ext sql $(name)
db.migrate:
	@migrate -path migrations -database postgres://localhost:5432/ecommerce?sslmode=disable up $(step)
db.rollback:
	@migrate -path migrations -database postgres://localhost:5432/ecommerce?sslmode=disable down $(step)
db.drop:
	@migrate -path migrations -database postgres://localhost:5432/ecommerce?sslmode=disable drop