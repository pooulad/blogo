# blogo📃
Complete blog web api written with golang/gin web server and postgres database

## Technology list in this project

 - JWT token
 - Godotenv
 - Postgres
 - GORM
 - Docker
 - Swagger
 - Gin


## How to run

### 🏳️First:

```bash
git clone https://github.com/pooulad/blogo
cd ./blogo
```

### ⚠️Dont forget to create .env file with required ```JWT_SECRET_TOKEN``` with jwt secret token value(you can generate from internet)

### 🐳Run with Docker:

1-Simple way with .sh file:

```bash
./run.sh
```

2-Run docker compose manually
```bash
docker compose up --build -d
```

### 🍉Run manaully:
1-In root of source you should run your go project with your db config in ./config/config.json or .env file
```bash
go run ./cmd/blogo --cfg ./config/config.json
```

2-Or you can add all flags one by one like this sample:
```bash
go run ./cmd/blogo --db_name="test" --db_port=5432 ...
```

#### All flags

| Flag name   | Description                  |
| ----------- | ---------------------------- |
| env         | env file address             |
| app_url     | application url              |
| port        | application port             |
| db_port     | database port                |
| db_name     | database name                |
| db_host     | database host                |
| db_username | database username            |
| db_password | database password            |
| db_sslmode  | database sslmode(true/false) |
| cfg         | confige file                 |

#### Sample config json file

```json
{
    "app_url": "0.0.0.0",
    "port": "8000",
    "environment": "development",
    "db": {
      "postgresql": {
        "port": "5432",
        "host": "postgres",
        "dbname": "test_blogo",
        "username": "postgres",
        "password": "postgres",
        "sslmode": "disable"
      }
    }
  }
  
```

#### All endpoints

you can see all of them in ./docs/insomnia directory with .json or .har or .yaml extention
```bash
cd ./docs/insomnia
``` 

#### TODO Checklist

This section tracks the progress of implemented features in blogo.

- [x] Dockerize
- [x] Add swagger
- [x] Releaser
- [ ] Validation for post methods
- [ ] Testing

## 🛡️ License

This project is licensed under the MIT License.

