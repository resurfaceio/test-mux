PROJECT_NAME=test-mux

start:
	@docker stop resurface || true
	@docker build -f Dockerfile -t test-mux --no-cache .
	@docker run -detach -p 4000:4000 --name resurface-mux -t test-mux
	# @docker-compose up --detach

stop:
	@docker stop resurface-mux
	@docker image rmi -f test-mux

restart:
	echo "restart script here"

test:
	echo "test script here"

clear:
	@docker system prune -a