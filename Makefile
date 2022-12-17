nodemon:
	nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run main.go

setup:
	docker pull mongo
	docker run -d --name mongodb-user-api -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=root mongo

run-all:
	make start-mongo
	make run-app

run-app:
	go run main.go

start-mongo:
	docker start mongodb-user-api

stop-mongo:
	docker stop mongodb-user-api