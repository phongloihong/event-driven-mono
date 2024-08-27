.PHONY: build-all build-app start logs-app logs-mongodb ssh-app ssh-mongo create-db proto

build-all:
	docker-compose build
build-app:
	docker-compose build app
start:
	docker-compose up -d && docker exec -it event-cart-bff bash -c "cd /app/services/cart-bff  && go work use -r ."
logs-app:
	docker logs -f event-cart-bff
ssh-app:
	docker exec -it event-cart-bff bash
test:
    gotestsum --format testname -- ./...