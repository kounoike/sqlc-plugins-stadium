all: sqlc-crud sqlc-fix-schema sqlc-name

sqlc-crud:
	cd plugin/sqlc-crud && go build .

# sqlc-crud.wasm:
# 	cd plugin/crud && tinygo build -o sqlc-crud.wasm -target=wasi main.go
# 	openssl sha256 plugin/sqlc-crud.wasm

sqlc-fix-schema:
	cd plugin/sqlc-fix-schema && go build .

sqlc-name:
	cd plugin/sqlc-name && go build .

install:
	cp plugin/sqlc-crud/sqlc-crud plugin/sqlc-fix-schema/sqlc-fix-schema plugin/sqlc-name/sqlc-name $(GOPATH)/bin/

run:
	rm -r sample/database/schema_fix && mkdir sample/database/schema_fix
	rm -r sample/database/queries_crud && mkdir sample/database/queries_crud
	cd sample && sqlc generate -f sqlc-fix.yaml && sqlc generate -f sqlc-crud.yaml && sqlc generate
