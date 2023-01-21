get-report:
	sudo pgbadger logs/*.json -J 4 --outdir reports/ -o report-$(date +'%Y-%m-%d-T-%H:%M:%S').html

clone-pgbadger:
	git clone https://github.com/darold/pgbadger	

install-pgbadger:clone-pgbadger
	cd pgbadger && perl Makefile.PL && make && sudo make install && cd .. && sudo rm -r pgbadger


set-up-user:
	sudo docker cp ./setUpServiceUser.sql postgres:/ && \
	sudo docker exec -it postgres psql -U ${DB_USER} -d base -a -f \
	./setUpServiceUser.sql -v user=${TEST_DB_USER} -v passwrd=${TEST_DB_PASSWORD} \
	-v dbname=${TEST_DB_NAME}

set-docker-user:
	@sudo docker exec -it postgres psql -U ${DB_USER} -d base -a -f \
	/docker-entrypoint/setUpServiceUser.sql \
	-v user=${TEST_DB_USER} -v passwrd=${TEST_DB_PASSWORD} \
	-v dbname=${TEST_DB_NAME}
