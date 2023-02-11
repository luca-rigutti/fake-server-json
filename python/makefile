make start_server:
	docker-compose -f docker-compose.yaml up -d

make start_server_rebuild:
	docker-compose -f docker-compose.yaml up --build

make be_bash:
	docker exec -it fake-server-json_fake-server-json_1 /bin/sh