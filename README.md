## Database Migration

Database migration dibuat menggunakan tools [golang-migrate](https://github.com/golang-migrate/migrate).\
Untuk menjalankan migration up gunakan:\
`migrate -path db/migration -database "mysql://root:secret@tcp(localhost:1357)/article" -verbose up`\
Untuk menjalankan migration down gunakan:\
`migrate -path db/migration -database "mysql://root:secret@tcp(localhost:1357)/article" -verbose down `

> Sesuaikan connection string database : mysql://**username**:**password**@tcp(**address**:**port**)/**database_name** \

## ENV Variable

Sebelum aplikasi dapat berjalan dibutuhkan file **app.env** dengan isi

```bash
DB_DRIVER=mysql
DB_USERNAME=root
DB_PASSWORD=secret
DB_ADDRESS=localhost:1357
SERVER_ADDRESS=localhost:1234
```

**SERVER_ADDRESS** adalah address dimana server (api) akan berjalan

> Sesuaikan konfigurasi DB_USERNAME, DB_PASSWORD, dan DB_ADDRESS

## Build and run

```bash
go run .
```
