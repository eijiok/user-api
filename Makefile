# installing go dependencies
install:
	go get
	go mod tidy

# this is a command to rerun every time one source file is saved
nodemon:
	nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run main.go

# makes a mongo setup
mongo-setup:
	docker pull mongo
	docker run -d --name mongodb-user-api -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=root mongo

# install swagger for Linux / MAC via docker
swagger-setup:
	docker pull quay.io/goswagger/swagger
	alias swagger='docker run --rm -it  --user $(id -u):$(id -g) -e GOCACHE=/tmp -e  GOPATH=$(go env GOPATH):/go -v $HOME:$HOME -w $(pwd) quay.io/goswagger/swagger'

swagger-docs:
	swagger generate spec -i ./swagger.yml -m -o ./swagger.json

# if you didn't started the mongodb container, you can run this command
run-all:
	@make start-mongo
	@make run-app

# runs the application. Make sure the mongo db is already running
run-app:
	go run main.go

# starts just the mongo db on docker
start-mongo:
	docker start mongodb-user-api

# stops mongo db on docker
stop-mongo:
	docker stop mongodb-user-api