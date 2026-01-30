OS := $(shell uname -s 2>/dev/null || echo Windows)

# Variables
ENVIRONMENT ?= dev
CONTAINER_NAME=${APP_NAME}-app-${ENVIRONMENT}
POSTGRES_CONTAINER_NAME=${APP_NAME}-db-${ENVIRONMENT}

dep: 
	go mod tidy

run: 
	go run main.go

watch:
	go run main.go --watch

seeder:
	go run main.go --seeder

migrate:
	go run main.go --migrate

both:
	go run main.go --migrate --seeder

init-dev:
	docker-compose -f docker-compose.dev.yml up --build -d

up-dev: 
	docker-compose -f docker-compose.dev.yml up -d

down-dev-volume:
	docker-compose -f docker-compose.dev.yml down -v 

down-dev:
	docker-compose -f docker-compose.dev.yml down

init-prod:
	docker-compose -f docker-compose.prod.yml up --build -d

up-prod:
	docker-compose -f docker-compose.prod.yml down
	docker-compose -f docker-compose.prod.yml up -d

down-prod-volume:
	docker-compose -f docker-compose.prod.yml down -v 

down-prod:
	docker-compose -f docker-compose.prod.yml down

logs-prod:
	docker-compose -f docker-compose.prod.yml logs -f app-prod

logs-dev:
	docker-compose -f docker-compose.dev.yml logs -f app-dev

build-prod:
	docker-compose -f docker-compose.prod.yml build app-prod

rebuild-prod: down-prod init-prod

rebuild-dev: 
	docker-compose -f docker-compose.dev.yml down
	docker-compose -f docker-compose.dev.yml up -d

migrate-docker:
	docker exec -it crs-backend-app-dev go run main.go --migrate

seeder-docker:
	docker exec -it crs-backend-app-dev go run main.go --seeder

both-docker:
	docker exec -it crs-backend-app-dev go run main.go --migrate --seeder

help:
	@echo "Usage: make [target]"
	@echo "Targets:"
	@echo "  dep                         Tidy dependencies"
	@echo "  run                         Run the application"
	@echo "  watch                       Run program with auto loading"
	@echo "  seeder                      Seed the database"
	@echo "  migrate                     Run database migrations"
	@echo "  both                        Run migrations and seeder"
	@echo ""
	@echo "Development:"
	@echo "  init-dev                    Initialize dev environment with build"
	@echo "  up-dev                      Start dev containers"
	@echo "  down-dev                    Stop dev containers"
	@echo "  down-dev-volume             Stop dev containers and remove volumes"
	@echo "  logs-dev                    View dev logs"
	@echo "  rebuild-dev                 Rebuild and restart dev"
	@echo ""
	@echo "Production:"
	@echo "  init-prod                   Initialize prod environment with build"
	@echo "  up-prod                     Start prod containers"
	@echo "  down-prod                   Stop prod containers"
	@echo "  down-prod-volume            Stop prod containers and remove volumes"
	@echo "  logs-prod                   View prod logs"
	@echo "  build-prod                  Build prod image"
	@echo "  rebuild-prod                Rebuild and restart prod (with downtime)"
	@echo ""
	@echo "  help                        Show this help message"