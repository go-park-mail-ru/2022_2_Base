#sudo docker exec -it postgres bash
printf "DROP TABLE:\n"
sudo docker exec -it postgres psql -U spuser -d base -a -f ./docker-entrypoint/drop.sql -S | grep -E "(NOTICE|ERROR)" && printf "\n"
printf "CREATE TABLE:\n"
sudo docker exec -it postgres psql -U spuser -d base -a -f ./docker-entrypoint/init.sql | grep -E "(NOTICE|ERROR)" && printf "\n"
printf "POPULTE TABLE:\n"
sudo docker exec -it postgres psql -U spuser -d base -a -f ./docker-entrypoint/populate.sql  | grep -E "(NOTICE|ERROR)" && printf "\n"
sudo docker exec -it postgres psql -U spuser --dbname=base
