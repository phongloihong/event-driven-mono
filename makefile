.PHONY: build-all start logs-app logs-mongodb ssh-app update-workspace mock

build-all:
	docker-compose build
start:
	docker-compose up -d 
logs:
	docker compose logs -f
mock:
	mockery --name=IRepository --dir=./libs/database/mongoLoader --outpkg=mocks --output=./libs/mocks \
    && mockery --name=ILogger --dir=./libs/log --outpkg=mocks --output=./libs/mocks