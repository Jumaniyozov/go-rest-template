include .env
CURRENT_DIR=$(shell pwd)

APP=$(shell basename ${CURRENT_DIR})
API_CMD_DIR=${CURRENT_DIR}/cmd/api

TAG=0.0.1
ENV_TAG=0.0.1

.PHONY: db/create/docker/postgres/container
db/create/docker/postgres/container:
	@echo 'Creating docker postgres container...'
	docker run -d --name ${DOCKER_POSTGRES_CONTAINER_NAME} -p ${POSTGRES_PORT}:5432 -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} -v ${DOCKER_VOLUME_NAME}:/var/lib/postgresql/data --restart always postgres

.PHONY: db/create/docker/database
db/create/docker/database:
	@echo 'Creating docker database...'
	docker exec -it ${DOCKER_POSTGRES_CONTAINER_NAME} createdb --username=${POSTGRES_USER} --owner=${POSTGRES_USER} ${POSTGRES_DB_NAME}

.PHONY: db/drop/docker/database
db/drop/docker/database:
	@echo 'Dropping docker database...'
	docker exec -it ${DOCKER_POSTGRES_CONTAINER_NAME} dropdb --username=${POSTGRES_USER} ${POSTGRES_DB_NAME}

# migrations/new name=$1: create a new database migration
.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=${CURRENT_DIR}/migrations ${name}

.PHONY: db/migrate/up
db/migrate/up: confirm
	@echo 'Running up migrations...'
	migrate -path ${CURRENT_DIR}/migrations -database 'postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB_NAME}?sslmode=disable' up

.PHONY: db/migrate/down
db/migrate/down: confirm
	@echo 'Running down migrations...'
	migrate -path ${CURRENT_DIR}/migrations -database 'postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB_NAME}?sslmode=disable' down

.PHONY: db/sqlc
db/sqlc:
	@echo 'Generating sqlc files...'
	sqlc generate

## vendor: tidy and vendor dependencies
.PHONY: vendor
vendor:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Vendoring dependencies...'
	go mod vendor

.PHONY: audit
audit: vendor
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	staticcheck ./...
	@echo 'Running tests...'
	go test -race -vet=off ./...

.PHONY: build/api
build/api:
	@echo 'Building cmd/api...'
	GOOS=linux GOARCH=amd64 go build -ldflags="-s" -o=${CURRENT_DIR}/bin/linux_amd64/api ${API_CMD_DIR}

.PHONY: run/api
run/api:
	@echo 'Starting a server...'
	go run ${API_CMD_DIR}/main.go


.PHONY: swag
swag:
	@echo 'Initializing swagger...'
	swag init -dir ./cmd/api/ -o ./api/docs


.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]


