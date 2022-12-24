make start_server:
	docker-compose -f docker-compose.yaml up -d

make start_server_rebuild:
	docker-compose -f docker-compose.yaml up -d --build