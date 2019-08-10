main: dev-build

dev:
	docker-compose up

dev-build:
	docker-compose up --build

stop:
	docker-compose down