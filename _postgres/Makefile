get-report:
	sudo pgbadger logs/*.json -J 4 --outdir reports/ -o $(date +'%Y-%m-%d-T-%H:%M:%S').html

clone-pgbadger:
	git clone https://github.com/darold/pgbadger	

install-pgbadger:clone-pgbadger
	cd pgbadger && perl Makefile.PL && make && sudo make install && cd .. && sudo rm -r pgbadger