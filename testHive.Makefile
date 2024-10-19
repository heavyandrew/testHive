up:
    docker-compose up --build

dev-up:
    docker-compose up -d db

down:
    docker-compose down

test:
    go test ./... -v

swagger:
    swag init -g cmd/main.go