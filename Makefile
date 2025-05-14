export PGPASSWORD=postgres

build:
	docker build -t db .

create:
	docker run -d \
  --name db \
  -e POSTGRES_PASSWORD=postgres \
  -p 5432:5432 \
  db

start:
	docker start db

stop:
	docker stop db

remove:
	docker rm db


reset:
	psql -h localhost -p 5432 -U postgres -d template1 -c "DROP DATABASE postgres WITH (FORCE);"
	psql -h localhost -p 5432 -U postgres -d template1 -c "CREATE DATABASE postgres;"
	psql -h localhost -p 5432 -U postgres -d postgres -f seeds.sql

seeds:
	psql -h localhost -p 5432 -U postgres -d postgres -f seeds.sql


run-api-server:
	cd employees-server && go run main.go