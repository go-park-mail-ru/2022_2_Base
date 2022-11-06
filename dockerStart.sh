# sudo docker-compose up -d
sudo docker-compose up -d --build
sudo docker container stop "$(sudo docker container ls | grep -Eo '[a-zA-Z0-9]{12}')"
sudo docker container prune -f
sudo docker image prune -f
docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' container_name_or_id
# sudo docker exec -it postgres psql -U user
# sudo docker container rm *

# sudo docker compose stop
# sudo docker compose up -d --build

# sudo docker exec -it postgres psql -U spuser \dt

# docker exec -it postgres bash
# psql -U postgres
# \c - connect to db

# psql -h localhost -U postgres