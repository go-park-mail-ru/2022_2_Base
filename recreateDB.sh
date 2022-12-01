#sudo docker exec -it postgres bash
#docker cp ./_postgres postgres:/
printf "Copy scripts:\n"
sudo docker cp ./_postgres postgres:/
printf "DROP TABLE:\n"
sudo docker exec -it postgres psql -U spuser -d base -a -f ./_postgres/drop.sql -S | grep -E "(NOTICE|ERROR)" && printf "\n"
printf "CREATE TABLE:\n"
sudo docker exec -it postgres psql -U spuser -d base -a -f ./_postgres/init.sql | grep -E "(NOTICE|ERROR)" && printf "\n"
printf "POPULTE TABLE:\n"
sudo docker exec -it postgres psql -U spuser -d base -a -f ./_postgres/populate.sql  | grep -E "(NOTICE|ERROR)" && printf "\n"
#sudo docker exec -it postgres psql -U spuser --dbname=base
