# ==================================================================================== #
# PREREQUIREMENTS
# ==================================================================================== #

bin-deps:
	@$(MAKE) dependency

.PHONY: dependency
dependency:
	go install github.com/not-for-prod/implgen@latest
	go install github.com/not-for-prod/clay/cmd/protoc-gen-goclay@latest
	go install github.com/matryer/moq@latest
	go install github.com/pressly/goose/v3/cmd/goose@latest

# ==================================================================================== #
# INFRASTRUCTURE
# ==================================================================================== #

.PHONY: infra ## поднимает инфрастуктуру для проекта
.SILENCE:
infra:
	docker-compose -f ./build/docker-compose.yaml up -d --build --force-recreate --wait

infra-stop:
	docker-compose -f ./build/docker-compose.yaml down

# ==================================================================================== #
# MIGRATIONS
# ==================================================================================== #

DB_NAME=gift-box
DB_USER=postgres
DB_PORT=5432
MIGRATION_FOLDER=./tools/migrations

.PHONY: migrations-up ## накатывает миграции на базу данных
migrations-up:
	$(LOCAL_BIN)/goose postgres 'host=localhost port=${DB_PORT} user=${DB_USER} sslmode=disable dbname=${DB_NAME}' -dir ${MIGRATION_FOLDER} -allow-missing up

.PHONY: migrations-reset ## накатывает миграции на базу данных
migrations-reset:
	$(LOCAL_BIN)/goose postgres 'host=localhost port=${DB_PORT} user=${DB_USER} sslmode=disable dbname=${DB_NAME}' -dir ${MIGRATION_FOLDER} -allow-missing reset

.PHONY: migrations ## накатывает миграции на базу данных
migrations: migrations-reset migrations-up

# ==================================================================================== #
# CODEGEN
# ==================================================================================== #

generate:
	buf dep update
	buf dep prune
	buf build
	buf generate

implgen:
	implgen --src internal/generated/pb/yelp/review/v1/service_grpc.pb.go --interface-name ReviewServiceServer --dst internal/delivery/api