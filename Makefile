PROJECT_NAME=test-mux

start:
	@docker stop resurface || true
	@docker build -t test-mux --no-cache .
	# @docker run -detach -p 4000:4000 --name resurface-mux -t test-mux
	@docker-compose up --detach

stop:
	@docker-compose stop
	@docker-compose down --volumes --remove-orphans
	@docker image rmi -f test-mux

	# @docker stop resurface-mux
	# @docker image rmi -f test-mux

restart:
	echo "restart script here"

test:
	echo "test script here"

clear:
	@docker system prune -a