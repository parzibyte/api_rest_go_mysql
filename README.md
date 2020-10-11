# api_rest_go_mysql
 Creación de una API REST con Golang/Go y MySQL

## Installation
* Clone this repo
* Install the dependencies with:
```
go get github.com/go-sql-driver/mysql
go get github.com/gorilla/mux
go get github.com/joho/godotenv
```
* Create a database in MySQL
* Create a file called **.env** based on **.env.example**
* Configure credentials in **.env**
* In the database (you can use phpmyadmin or console) create
the tables according to **scheme.sql**
* Compile with `go build`
* Execute the exe, normally it will be **api_rest_go_mysql.exe** 
(if Windows firewall ask, give it permission, no problem)
    * Optionally if you ask me how I compile on Windows and format my code,
the command is `gofmt -w . && go build && api_rest_go_mysql`
* Now you can test the API with postman or anything else. 
If you didn't make changes, it will be on
localhost:8000

## About
Created by Parzibyte (https://parzibyte.me/blog)
