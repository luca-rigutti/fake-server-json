docker-compose = docker-compose 
fileDockerCompose = docker-compose.yaml

make start_server:
	$(docker-compose) -f $(fileDockerCompose) up -d

make start_server_rebuild:
	$(docker-compose) -f $(fileDockerCompose) up -d --build

make down:
	$(docker-compose) -f $(fileDockerCompose) down

make be_bash:
	docker exec -it fake-server-json_fake-server-json_1 /bin/sh