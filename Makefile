run:
	go run github.com/atkinsonbg/go-gmux-proper-unit-testing

up:
	GOOS=linux GOARCH=amd64 go build -o go-gmux-proper-unit-testing-api
	docker-compose down
	COMPOSE_DOCKER_CLI_BUILD=1 docker-compose up --build --detach
	docker-compose logs -f

down:
	rm -rf go-gmux-proper-unit-testing-api
	docker-compose down