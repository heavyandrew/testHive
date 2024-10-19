# Загрузка переменных из .env
include .env
export $(shell sed 's/=.*//' .env)

up:
	docker-compose up -d

dev-up:
	docker-compose up -d db

migrate:
	docker-compose exec db psql -U $(DB_USER) -d $(DB_NAME) -f /migrations/001_create_tables.sql

down:
	docker-compose down

test:
	go test ./... -v

swagger:
	swag init -g cmd/main.go