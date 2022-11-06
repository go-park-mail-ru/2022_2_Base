sudo -i
#   
docker-compose up -d --build
# docker exec -it postgres psql -U spuser
# docker container rm *
# docker logs -f go
# Show tables: \dt
# Show databases: \l
# psql -U spuser --dbname=base
# docker exec -it postgres bash
# psql -U spuser --dbname=base -a -f init.sql


#docker-compose --env-file ./.env  up -d --build
# docker exec -it postgres bash
#psql -U spuser --dbname=base