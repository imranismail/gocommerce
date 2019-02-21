
serve:
	@realize start
gen.migration:
	@migrate create -dir migrations -ext sql $(name)
db.migrate:
	@migrate -path migrations -database postgres://localhost:5432/ecommerce?sslmode=disable up $(step)
db.force:
	@migrate -path migrations -database postgres://localhost:5432/ecommerce?sslmode=disable force $(version)
db.rollback:
	@migrate -path migrations -database postgres://localhost:5432/ecommerce?sslmode=disable down $(step)
db.drop:
	@migrate -path migrations -database postgres://localhost:5432/ecommerce?sslmode=disable drop