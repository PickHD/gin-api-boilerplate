# GIN API Boilerplates

## Pre-requisites :
- **PostgreSQL** must be installed in your machines.

## Setup :
- Copy ```.env.sample``` into ```.env```, (fill the env with your own configuration).
- Run the api with : ```go run .```.
- Server will running on ```http://localhost:{your_port}```. 

## Features :
- Well structured boilerplates (can be scalling into microservices or monolith).
- Database (**postgres**) configuration (including AutoMigrate from **gorm** ORM).
- Router setup,api versioning **(v1)**, with bonus ```pinging server``` endpoint.
- Documentation Templates (with **OA3 Swagger**).
- Middlewares **CORS** & **RequestId** Header.
- Bunch of utilites including : **response formatter,generator,paginator,validator**.
- Included with sample of create apps (**base,models,controllers,services**).