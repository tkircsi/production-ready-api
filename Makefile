

server:
	DB_USERNAME=postgres \
	DB_PASSWORD=example \
	DB_HOST=localhost \
	DB_PORT=5432 \
	DB_TABLE=postgres \
	go run cmd/server/main.go 


.PHONY: server