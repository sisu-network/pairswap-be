dev-up:
	mkdir -p docker/db
	@docker compose \
		-f docker/docker-compose.yml up -d

dev-down:
	@docker compose \
		-f docker/docker-compose.yml down
