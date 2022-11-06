make build:
	docker-compose up -d --build

make up:
	docker-compose up -d

make container:rm:
	docker container rm *

make container:ls:
	docker container ls

make lol:
	ls	

make kek:
	echo "asd"