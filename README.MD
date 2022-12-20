# REST API in Go Lang

This project is an example of a CRUD Rest API built in Go Lang. 

## Setup

This project was made in Go Lang with the database Mongodb.

### MongoDB
#### New Configuration on docker
If you don't have a Mongodb, but you have a docker installed in your machine, run the following commands:

```sh
docker pull mongo
docker run -d --name mongodb-user-api -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=root mongo
```
 or, alternatively, you can run
 ```sh
make mongo-setup
```

#### Using an existing mongoDB instalation
You can connect to an existing running mongoDB editing the .env file

#### Commands to start / stop Mongodb on Docker
To start the mongodb container:
 ```sh
docker start mongodb-user-api
```
To stop the mongodb container:
```sh
docker stop mongodb-user-api
```

### Running the application
Install dependencies:
```sh
make install
```

Command to start running the application
```sh
go run main.go
```

### User API

The user API can be accessed on:

    http://localhost:8080/docs



#### Creating a user
Create a user.
The fields: name, email and password are required
It is a not secured endpoint to allow users to register to the application.
`POST /api/users`
```sh
    curl --request POST \
        -d '{"name":"Eiji", "password":"asdfqwer", "address":"rua xpto, 123", "birthday": "1981-04-14T07:01:02.000Z", "email":"eijiok@gmail.com"}' \
        --url http://localhost:8080/api/users \
        --header 'Content-Type: application/json' 
``` 


#### List users 

By now it has to have the header with Authorization parameter set with anything. 
It will change after JWT implementation. 

`GET /api/users`
```sh
    curl --request GET \
        --url http://localhost:8080/api/users \
        --header 'Authorization: user-authorization-header' \
        --header 'Content-Type: application/json' 
``` 

#### Read a user

It has to have the header with Authorization parameter.

`GET /api/users/{id}`
```sh
    curl --request GET \
        --url http://localhost:8080/api/users/639e4dd6892ab2a78e806f94 \
        --header 'Authorization: user-authorization-header' \
        --header 'Content-Type: application/json' 
``` 

#### Update a user

It has to have the header with Authorization parameter.

`PUT /api/users/{id}` 
```sh
    curl --request PUT \
        -d '{"name":"Marcos"}' \
        --url http://localhost:8080/api/users/639e4dd6892ab2a78e806f94 \
        --header 'Authorization: user-authorization-header' \
        --header 'Content-Type: application/json' 
``` 


#### Delete a user
`DELETE /api/users/{id}`
```sh
    curl --request DELETE \
        --url http://localhost:8080/api/users/639e4dd6892ab2a78e806f94 \
        --header 'Authorization: user-authorization-header' \
        --header 'Content-Type: application/json' 
``` 

### TO DO:

- Implement a JWT authentication using the saved email/password
- Increase coverage tests
- Improve the generated swagger document
- Learn more about go libraries and refactor the code