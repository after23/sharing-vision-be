createdb:
	docker run --name mysql-latest -e MYSQL_ROOT_PASSWORD=secret -e MYSQL_DATABASE=article -p 1357:3306 -d mysql:latest

migrateup:
	migrate -path db/migration -database "mysql://root:secret@tcp(localhost:1357)/article" -verbose up

githubmigrateup:
	migrate -path db/migration -database "mysql://root:root@tcp(127.0.0.1:3306)/article" -verbose up

migratedown:
	migrate -path db/migration -database "mysql://root:secret@tcp(localhost:1357)/article" -verbose down 

sqlc:
	docker run --rm -v "D:\Legacy Documents\go-project\sharing-vision-be:/src" -w /src kjconroy/sqlc generate

test:
	go test -v -cover ./...

.PHONY: createdb migrateup migratedown sqlc test githubmigrateup