PROTO_DIR=proto
PROTO_FILES=$(shell find $(PROTO_DIR) -name "*.proto")

protoc:
	protoc -I=$(PROTO_DIR) --go_out=. --go-grpc_out=. $(PROTO_FILES)
migrateup:
	migrate -path internal/repository/db/migration -database "postgresql://user:userpass@localhost:5433/user_service?sslmode=disable" -verbose up
migrateup1:
	migrate -path internal/repository/db/migration -database "postgresql://user:userpass@localhost:5433/user_service?sslmode=disable" -verbose up 1
migratedown:
	migrate -path internal/repository/db/migration -database "postgresql://user:userpass@localhost:5433/user_service?sslmode=disable" -verbose down
migratedown1:
	migrate -path internal/repository/db/migration -database "postgresql://user:userpass@localhost:5433/user_service?sslmode=disable" -verbose down 1
new_migration:
	migrate create -ext sql -dir internal/repository/db/migration -seq $(name)
sqlc:
	sqlc generate


.PHONY: protoc migrateup migrateup1 migratedown migratedown1 new_migration sqlc