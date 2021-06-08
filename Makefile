PROJECT_NAME=test-gorilla/mux

start:
	@docker stop resurface || true
	@docker build -t test-mux --no-cache .
	# @docker start test-mux
	# @docker-compose up --detach

stop:
	echo "stop script here"

restart:
	echo "restart script here"

test:
	echo "test script here"

clear:
	@docker system prune -a