all: sqlc-crud sqlc-fix-schema

sqlc-crud:
	cd plugin/crud && go build -o $(GOPATH)/bin/sqlc-crud ./main.go

# sqlc-crud.wasm:
# 	cd plugin/crud && tinygo build -o sqlc-crud.wasm -target=wasi main.go
# 	openssl sha256 plugin/sqlc-crud.wasm

sqlc-fix-schema:
	cd plugin/fix-schema && go build -o $(GOPATH)/bin/sqlc-fix-schema ./main.go

run:
	rm sample/database/schema_fix/* sample/database/queries_crud/*
	cd sample && sqlc generate -f sqlc-fix.yaml && sqlc generate -f sqlc-crud.yaml && sqlc generate
