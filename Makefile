.PHONY: docs build

docs:
	swag init -o ./pkg/docs

dev:
	go run main.go

deploy:
	docker-compose build; docker-compose up -d

start:
	docker-compose up -d

restart:
	docker-compose restart

mup:
	docker exec -it soda soda --config "/project/first/database.yml" --path="/project/first/migrations" migrate up

mdown:
	docker exec -it soda soda --config "/project/first/database.yml" --path="/project/first/migrations" migrate down