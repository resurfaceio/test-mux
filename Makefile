PROJECT_NAME=test-mux

start:
	@docker stop resurface || true
	@docker build -t test-mux --no-cache .
	@docker-compose up --detach

stop:
	@docker-compose stop
	@docker-compose down --volumes --remove-orphans
	@docker image rmi -f test-mux

ping:
	@curl "http://localhost/ping"

restart:
	@docker-compose stop
	@docker-compose up --detach

test:
	echo "test script cmd goes here"

clear:
	@docker system prune -a