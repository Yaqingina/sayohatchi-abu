run:
	@zsh -c "set -a && source .env && go run cmd/main.go"

compose_up:
	@docker compose --project-name sayohatchi-abu -f docker-compose.yml up -d --build

compose_down:
	@docker compose --project-name sayohatchi-abu -f docker-compose.yml down

