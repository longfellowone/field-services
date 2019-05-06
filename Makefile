up:
	@echo ""
	@echo "Starting docker containers..."
	@echo ""
	docker-compose up -d
	@echo ""
	@echo "Postgres running on localhost:5432"
	@echo "Adminer running on localhost:8081"
	@echo "gRPC service running on localhost:9090"
	@echo "Envoy incoming on localhost:8080"
	@echo "Envoy outgoing on localhost:9090"

down:
	@echo ""
	@echo "Stopping docker containers..."
	@echo ""
	docker-compose down

build:
	@echo ""
	@echo "Rebuilding docker containers..."
	@echo ""
	docker-compose build

info:
	@echo ""
	docker ps
	@echo ""

logs:
	docker-compose logs -f

clean: down
	@echo ""
	@echo "Cleaning up..."
	@echo ""
	docker system prune
	docker volume prune