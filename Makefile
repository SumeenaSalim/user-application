up:
	docker compose up
build:
	docker compose up --build
down:
	docker compose down
migrate:
	docker compose -f docker-compose.flyway.yaml up