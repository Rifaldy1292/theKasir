.PHONY: db-up db-down dev-fe dev-be

# Start PostgreSQL database
db-up:
	docker-compose up -d

# Stop PostgreSQL database
db-down:
	docker-compose down

# Run Frontend (Nuxt)
dev-fe:
	cd frontend && npm run dev

# Run Backend (Go Gin)
dev-be:
	cd backend && go run cmd/api/main.go
