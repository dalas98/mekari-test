# Sample API Project 
#### by Yusuf Farhan Hasbullah

### Introduction
- This project use GIN as the framework. ([GIN](https://github.com/gin-gonic/gin)), PostgreSQL, Gorm and Goose
- Unit test on db repositories
- Auto db migrations

### Installation
1. Rename `.env.example` file to `.env`
2. Make sure database config in `.env` file
3. Create Database with name `api_go`

### How to Run
1. run
    ```
    docker-compose up -d
    ```
2. stop
    ```
    docker-compose down
    ```
3. if you need to re-build docker / update docker container (if you have any changes of code), please run 
    ```
    docker-compose up --build
    ```
