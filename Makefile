mysql:
	docker run --name mysql_8.0.30 -p 3306:3306 -e MYSQL_ROOT_PASSWORD=53cr3t01 -e MYSQL_DATABASE=banking -d mysql:8.0.30-debian

migrateup:
	migrate -path db/migration -database "mysql://root:53cr3t01@tcp(localhost:3306)/banking?tls=false" -verbose up

migrateup1:
	migrate -path db/migration -database "mysql://root:53cr3t01@tcp(localhost:3306)/banking?tls=false" -verbose up 1

migratedown:
	migrate -path db/migration -database "mysql://root:53cr3t01@tcp(localhost:3306)/banking?tls=false" -verbose down

migratedown1:
	migrate -path db/migration -database "mysql://root:53cr3t01@tcp(localhost:3306)/banking?tls=false" -verbose down 1

run:
	SERVER_ADDRESS=localhost SERVER_PORT=8000 DB_USER=root DB_PASSWD=53cr3t01 DB_ADDRESS=localhost DB_PORT=3306 DB_NAME=banking go run main.go

.PHONY: mysql migrateup migratedown run