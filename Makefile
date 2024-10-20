# Загрузка переменных из .env
include .env
export $(shell sed 's/=.*//' .env)

up:
	docker-compose up -d --build

dev-up:
	docker-compose up -d db

restart:
	docker-compose restart

build:
	docker-compose up -d db
	sleep 5
	docker-compose exec db psql -U $(DB_USER) -d $(DB_NAME) -f /migrations/001_create_tables.sql
	sleep 5
	docker-compose up -d --build

rebuild:
	docker-compose down
	docker volume rm -f testhive_db_data
	docker images rm --all
	docker-compose up -d db
	sleep 5
	docker-compose exec db psql -U $(DB_USER) -d $(DB_NAME) -f /migrations/001_create_tables.sql
	sleep 5
	docker-compose up -d --build

migrate:
	docker-compose exec db psql -U $(DB_USER) -d $(DB_NAME) -f /migrations/001_create_tables.sql

down:
	docker-compose down

test-reg:
	go test -count=1 ./testing/functional/register

test-login:
	go test -count=1 ./testing/functional/login

test-asset:
	go test -count=1 ./testing/functional/create-get-delete_asset

swagger:
	swag init -g cmd/main.go