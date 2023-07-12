createdb:
	docker run --name mysql-latest -e MYSQL_ROOT_PASSWORD=secret -e MYSQL_DATABASE=article -p 1357:3306 -d mysql:latest

migrateup:
	migrate -path db/migration -database "mysql://root:secret@tcp(localhost:1357)/article" -verbose up

migratedown:
	migrate -path db/migration -database "mysql://root:secret@tcp(localhost:1357)/article" -verbose down 

sqlc:
	docker run --rm -v "D:\Legacy Documents\go-project\sharing-vision-be:/src" -w /src kjconroy/sqlc generate

.PHONY: createdb migrateup migratedown sqlc