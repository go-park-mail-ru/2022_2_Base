test:
	go test -race -coverpkg=./... -coverprofile cover.out.tmp ./...; cat cover.out.tmp | grep -v "_easyjson.go" > cover1.out.tmp; cat cover1.out.tmp | grep -v ".pb.go" > cover2.out.tmp; cat cover2.out.tmp | grep -v "_mock.go" > cover.out; go tool cover -func cover.out

build:
	sudo docker-compose up -d --build && make set-service-user
build-it:
	sudo docker-compose up --build && make set-service-user

stop:docker-fix
	sudo docker-compose down

container-prune:
	sudo docker container prune -f

image-prune:
	sudo docker image prune -f

inspect-postgres:
	docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' postgres

postgres-bash:
	sudo docker exec -it postgres bash

docker-prune-all:
	sudo docker system prune -a

docker-fix:
	- sudo killall containerd-shim

connect-psql:
	sudo docker exec -it postgres psql -U spuser -d base

edit-psql-conf:docker-postgres-bash
	vim /var/lib/postgresql/data/postgresql.conf

set-env:
	set -a && source .env && set +a

set-service-user:
	 cd _postgres && make set-up-user
