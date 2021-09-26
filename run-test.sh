#!/bin/bash
echo "Running local server"
export PORT=80
go run server.go &
echo "Sleeping for 5 seconds"
sleep 5
echo "Setting dev env"
export IS_DEV=True
echo "Running test.."
test-resurface --request "go" --app-name "mux-test-app"

sleep 5
kill $!
echo "Done"
