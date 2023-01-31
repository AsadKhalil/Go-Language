# 

<h1 align="center"> Simple REST API </h1>



<div align= "center">
  <h4> REST API Created with Postgres Database</h4>
</div>

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
![GO](https://img.shields.io/badge/go-v1.9-lightgrey)
[![Forks](https://img.shields.io/github/forks/AsadKhalil/Go-Language.svg?logo=github)](https://github.com/AsadKhalil/Go-Language/network/members)
[![Stargazers](https://img.shields.io/github/stars/AsadKhalil/Go-Language.svg?logo=github)](https://github.com/AsadKhalil/Go-Language/stargazers)
[![Issues](https://img.shields.io/github/issues/AsadKhalil/Go-Language.svg?logo=github)](https://github.com/AsadKhalil/Go-Language/Face-Mask-Detection/issues)
[![LinkedIn](https://img.shields.io/badge/-LinkedIn-black.svg?style=flat-square&logo=linkedin&colorB=555)](https://www.linkedin.com/in/muhammad-asad10/)


&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;





## :warning: Tech/framework used

- Go
- PostGres
- Docker
- Rest API
- JWT Authentication



## 🚀&nbsp; Installation
1. Clone the repo
```
$ git clone https://github.com/AsadKhalil/Go-Language.git
```

2. Initialize Using
```
$ go mod init
```

3. Using Docker Compose to Start Data Base
```
$ docker-compose up
```

4. Build Using
```
$ go build .
```

3. Run Using
```
$ go run main.go
```


## Without Docker Set-up Database Using
- Install PostgreSQL on your system if it is not already installed.
- Open a terminal and connect to the PostgreSQL server using the psql command:
```
psql -U postgres
```
- Create a new database and user with the following commands:
```
CREATE DATABASE database;
CREATE USER username WITH PASSWORD 'password';
GRANT ALL PRIVILEGES ON DATABASE database TO username
```
Replace database and username with the desired names for your database and user.

- Connect to the new database using the psql command:
```
psql -U username -d database
```
- Create the tables Using init.sql file




## 💉:syringe: Testing on Post Man
POST MAN REQUESTS

```
POST 
http://localhost:8080/products
{
	"name": "My product",
	"description": "A product",
	"price": 100
}
```
PUT
```
{
	"id": 1,
	"name": "Updated product",
	"description": "An updated product",
	"price": 200
}
```
DELETE
```
http://localhost:8080/products
{
	"id": 1
}
```
GET
```
http://localhost:8080/products
```
```
http://localhost:8080/signup
signup

{
	"username": "johndoe",
	"email": "johndoe@example.com",
	"password": "password"
}
```
/login
{
	"username": "johndoe",
	"password": "password"
}
```
```
/credit-cards

{
	"card_number": "4111111111111111",
	"expiry_month": "01",
	"expiry_year": "2025",
	"cvv": "123",
	"name_on_card": "user"
}
```
```


## :clap: And it's done!
Feel free to mail me for any doubts/query 
:email: asadkhalil1@gmail.com

## :handshake: Contribution
Feel free to **file a new issue** with a respective title and description on the the [Go-Language](https://github.com/AsadKhalil/Go-Language) repository. If you already found a solution to your problem, **I would love to review your pull request**! 

## :heart: Owner
Made with :heart:&nbsp;  by [Asad](https://github.com/AsadKhalil)


## :eyes: License
MIT © [Asad]()

