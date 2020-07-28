run:
	go run github.com/atkinsonbg/go-gmux-proper-unit-testing

up:
	docker-compose down
	COMPOSE_DOCKER_CLI_BUILD=1 docker-compose up --build --detach
	docker-compose logs -f

down:
	rm -rf go-gmux-proper-unit-testing-api
	docker-compose down

docker:
	docker build -t github.com/atkinsonbg/go-gmux-proper-unit-testing/api:latest .

dockerrun:
	docker run -it github.com/atkinsonbg/go-gmux-proper-unit-testing/api:latest