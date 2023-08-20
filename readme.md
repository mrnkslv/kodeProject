# notes-app

* [Introduction](#introduction)
* [Quick_start](#quick_start)
* [Documentation](#documentation)
* [Dependencies](#dependencies)

## Introduction 

Notes-app API has the following functionality: 
- Register in the application
- Log in to the application (if you have an account)
- Create note (all the mistakes corrects automatically)
- Get all the user's notes 

## Quick_start
For quick start enter in the terminal command:
```
docker-compose up --build
```
In case of a database initialization error when launching the app container after the build, restart the container via docker commands or stop the process using (ctrl+c) and enter the command:
```
docker-compose up
```
To stop containers and remove containers,networks,volumes and images created by up: 
```
docker-compose down
```


## Documentation

Access URL examples and full documentation on REST API [here](documentation.md).

## Dependencies

### build
Docker, docker-compose

### database

PostgreSQL

### error correction service

Yandex Speller API
