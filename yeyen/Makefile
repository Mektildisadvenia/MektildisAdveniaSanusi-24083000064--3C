.PHONY: dev staging migrate-up migrate-down build

dev:
	docker-compose up --build

staging:
	docker-compose -f docker-compose.staging.yml up --build -d

down:
	docker-compose down
	docker-compose -f docker-compose.staging.yml down

# Run this inside the backend folder for local migrations
migrate-up:
	golang-migrate -path backend/migrations -database "postgres://postgres:postgres@localhost:5432/yeyen_dev?sslmode=disable" up

migrate-down:
	golang-migrate -path backend/migrations -database "postgres://postgres:postgres@localhost:5432/yeyen_dev?sslmode=disable" down

install:
	cd frontend && npm install
	cd backend && go mod tidy
