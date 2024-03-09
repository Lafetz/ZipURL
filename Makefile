run:
	go run ./cmd/api/main.go
# .PHONY: db/migrations/new
# db/migrations/new:
# 	@echo 'Creating migration files for ${name}...'
# 	migrate create -seq -ext=.sql -dir=./migrations ${name}
# .PHONY: db/migrations/up
# db/migrations/up: confirm
# 	@echo 'Running up migrations...'
# 	migrate -path ./migrations -database ${DB_DSN} up