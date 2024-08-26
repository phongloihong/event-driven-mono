.PHONY: build-all start logs-app logs-mongodb ssh-app update-workspace

build-all:
	docker-compose build
start:
	docker-compose up -d 
logs:
	docker compose logs -f