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
	@curl "http://localhost:8080/ping"

restart:
	echo "restart script here"

test:
	echo "test script here"

clear:
	@docker system prune -a