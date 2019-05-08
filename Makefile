up:
	@echo ""
	@echo "Starting docker containers..."
	@echo ""
	docker-compose up -d
	@echo ""
	@echo ":PORT SERVICE"
	@echo ""
	@echo ":5432 Postgres"
	@echo ":8081 Adminer"
	@echo ":8080 Envoy incoming"
	@echo ":9090 Envoy outgoing"
	@echo ":9090 gRPC Go server"
	@echo ""

down:
	@echo ""
	@echo "Stopping docker containers..."
	@echo ""
	docker-compose down
	@echo ""

build:
	@echo ""
	@echo "Rebuilding docker containers..."
	@echo ""
	docker-compose build
	@echo ""

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
	@echo ""