# Profile API - Go
### Final Project - BTPN Syariah Fullstack Developer Project Based Internship Program

## Description
This is a simple REST API for creating profiles, which is created for the final project of the BTPN Syariah Project Based Intership Program, In this API users can create new profiles using username, email, and password, users can also post new picture by using image url and add a title and caption to it.

## Installation
Clone the repository to your local machine:
```
git clone https://github.com/erwinmareto/task-5-pbi-btpns-erwin_mareto_wikas.git
```
Change into the project directory:
```
cd task-5-pbi-btpns-erwin_mareto_wikas
```
Install dependencies:

Please note that **gorm.io/driver/_postgres_** is for postgreSQL, please adjust to your machine
```
go get -u gorm.io/gorm gorm.io/driver/postgres github.com/gin-gonic/gin golang.org/x/crypto github.com/golang-jwt/jwt/v4 github.com/joho/godotenv github.com/githubnemo/CompileDaemon github.com/asaskevich/govalidator
```
Create a ```.env``` file with the following fields with approriate values for your environment:

```javascript
PORT=8000
DB="host={YOUR_DB_HOST} user={YOUR_DB_USERNAME} password={YOUR_DB_PASSWORD} dbname={NAME_OF_DB} port={YOUR_DB_PORT} sslmode=disable"
JWT_SECRET=any_secret_key
```

Run the project:
```
go run main.go
```

OR

If you are developing:
```
compiledaemon -command="./profile-api-go"
```