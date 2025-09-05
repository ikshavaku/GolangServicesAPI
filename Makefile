DB_URL=postgres://${db_user}:${db_password}@${db_server}:${db_port}/${db_name}?sslmode=disable
MIGRATIONS_DIR=./migrations

build:
	. ./local.envrc
	echo $(SERVER_PORT)
	go mod download
	go generate -tags wireinject ./...
	CGO_ENABLED=0 GOARCH=amd64 go build -o ./target/build ./cmd/app/...

goose-up:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" up

goose-down:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" down

goose-create:
	goose -dir $(MIGRATIONS_DIR) create $(name) sql

run-server:
	make build
	./target/build

test:
	go test -v -cover -coverpkg=./... ./... -coverprofile=tests/coverage.out -v