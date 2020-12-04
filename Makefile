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

dockertest:
	docker build -f Dockerfile.test -t github.com/atkinsonbg/go-gmux-proper-unit-testing/tests:latest .

dockertestrun:
	docker run -it github.com/atkinsonbg/go-gmux-proper-unit-testing/tests:latest

testlocal:	dockertest
	docker run -it -e GIT_URL='' github.com/atkinsonbg/go-gmux-proper-unit-testing/tests:latest

testremote:
	docker run -it -e GIT_URL='https://github.com/atkinsonbg/go-gmux-unit-testing.git' github.com/atkinsonbg/go-gmux-proper-unit-testing/tests:latest