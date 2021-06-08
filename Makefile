PROJECT_NAME=test-gorilla/mux

start:
	@sudo docker stop resurface || true
	@sudo docker build -f Dockerfile -t test-mux --no-cache .
	@sudo docker run -detach -p 4000:4000 -t test-mux
	# @docker-compose up --detach

stop:
	echo "stop script here"

restart:
	echo "restart script here"

test:
	echo "test script here"

clean:
	@sudo docker system prune -a