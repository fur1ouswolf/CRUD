.PHONY: run stop shutdown swagger

run:
	docker compose --env-file .env up -d --build

stop:
	docker compose --env-file .env down

shutdown:
	docker compose --env-file .env down --volumes

swagger:
	swag init --parseDependency --parseInternal --parseDepth 1 -g cmd/app/main.go -o docs/
	swag fmt