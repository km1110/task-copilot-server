init:
	docker compose build --no-cache

up:
	docker compose up

down:
	docker compose down

database:
	docker compose exec -it postgresql bash