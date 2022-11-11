make build:
	sudo docker-compose up -d --build

make build:d:
	sudo docker-compose up --build

make stop:
	sudo docker container stop "$(sudo docker container ls | grep -Eo '[a-zA-Z0-9]{12}')"

make container:prune:
	sudo docker container prune -f

make image:prune:
	sudo docker image prune -f

make inspect:postgres:
	docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' postgres

make docker:postgres-bash:
	sudo docker exec -it postgres bash

make docker:prune-all:
	sudo docker system prune -a
	