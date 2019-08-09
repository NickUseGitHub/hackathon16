main: dev

dev:
	docker-compose up

dev-build:
	docker-compose up --build

stop:
	docker-compose down