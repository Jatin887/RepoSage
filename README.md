# Balkanid Summer Internship Task

Backend Assignment | BalkanID

## About the Backend

This project aims to create a backend system that first authorize the user using Github Oauth and than can fetch the repositories (personal and organization) data from the Github API, normalize and deduplicate the data, store it in a Postgres database, and convert it to CSV format before providing an API endpoint for users to download the CSV file.

## Key Features
- OAuth using Github
- Mulitilevel logging 
- Network retries in case of failure

## Tech Stack
- Go
- Gin
- Gorm
- PostgreSQL
- Docker

## Major Packages used in this project

- **gin**: Used Gin in my project for its high performance and efficient handling of high traffic, while also providing a flexible and easy-to-use API for building web applications in Go.
- **viper**: Used for loading configuration from the `.env` file. 
- **gorm**: Used for mapping the Go structs to database tables and vice versa.
- **logrus**:Used for multilevel logging in the project and used to create a log file.


- Check more packages in `go.mod`.


### Getting Started

To get started:

- Clone the repo.

```shell
git clone https://github.com/BalkanID-University/balkanid-summer-internship-vit-vellore-2023-Jatin887
```

- Change into the directory.

```shell
cd balkanid-summer-internship-vit-vellore-2023-Jatin887
```

### Run

#### Environment Variables

```shell
touch .env
```

#### Run without Docker
```shell
- Create a file `.env` similar to `.env.example` at the root directory with your configuration.
- Install `go` if not installed on your machine.
- Install `PostgresSQl` if not installed on your machine.
- Access API using `http://localhost:3000`
```

#### Run with Docker
```shell
- Create a file `.env` similar to `.env.example` at the root directory with your configuration.
- Install Docker and Docker Compose.
- Build `docker compose up --build`
- Run `docker-compose up -d`.
- Access API using `http://localhost:300`
```
### Instructions
```shell
- Step1 - Open your browser and type localhost::3000
- Step2 - Press on login button 
- Step3 - Enter Credentials
- Step4 - Copy Access Token
- Step5 - See the postman api documentaion to get  the output which is attached below
```


### The Complete Project Folder Structure

```
.
├── Dockerfile
├── api
│   ├── controller
│   │   ├── auth_controller.go
│   │   ├── dashboard_controller.go
│   │   ├── download_controller_test.go
│   │   └── logout_controller.go
│   └── middleware
│       └── middelware.go
├── initializers
│   ├── app.go
│   ├── database.go   
│   ├── logger_init.go
│   └── load_env_variables.go
├── cmd
│   └── main.go
├── docker-compose.yaml
├── models
│   ├── github_acess_token_response.go
│   ├── github_user_data_model.go
│   ├── license_model.go
│   ├── org_repo_model.go
│   ├── org_model.go
│   ├── owner_model.go
│   ├── permission_model.go
│   ├── repo_model.go
│   ├── repo_response.go
│   ├── user_model.go
│   └── user_repo_model.go
├── go.mod
├── go.sum
├── config
│   └── config.go
├── migrate
│   └── migrate.go
├── repository
│   ├── task_repository.go
│   ├── user_repository.go
│   └── user_repository_test.go
└── helpers
    ├── convert_data.go
    ├── filter_save_data.go
    ├── get_clientId.go
    ├── get_data_database.go
    ├── git_data.go
    ├── request_retry.go
    └── save_user.go
```

### API documentation of BalkanID Task

<a href="https://documenter.getpostman.com/view/17031765/2s93XsZSUp" target="_blank">
    <img alt="View API Doc Button" src="https://github.com/amitshekhariitbhu/go-backend-clean-architecture/blob/main/assets/button-view-api-docs.png?raw=true" width="200" height="60"/>
</a>

<p align="center">Made with ❤ by Jatin Fulwani</p>