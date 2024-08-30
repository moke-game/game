# EverJoy Game Servers

## Introduction

This project is EverJoy Game Servers

## Architecture

![architecture](./draws/game.drawio.png)

## How to run

* deploy infrastructure:
  ```shell
   # add ./deployment/docker-compose/.env file to custom your environment if you have
   docker compose -f ./deployment/docker-compose/infrastructure.yaml up -d
  ```

* run service:
  ```shell
  # fix the game-name to your game name 
    go run ./cmd/{game-name}/service/main.go
  ```

## How to build docker image?

```shell
# replace <your_register_url> to your register url,eg: game0.registry.com
# replace <your_server_name> to your server name,eg: game0,gm
# replace <your_git_access_token> to your git access token
# how to get your git access token?: https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens#creating-a-personal-access-token-classic
docker buildx build -t <your_register_url>:latest --build-arg APP_NAME=<your_server_name> --build-arg GIT_PWD=<your_git_access_token>  -f .\build\package\docker\Dockerfile . --push
```

## How to test

### Integration Test

* build your interactive client:
   ```shell
     go build -o {game-name}.exe ./cmd/{game-name}/client/main.go 
   ```
* run your interactive client:
    ```shell
     # help
     ./{game-name}.exe shell
    ```
  tips: http client use Postman to connect `localhost:8081`.

### Load Test

* install [k6](https://grafana.com/docs/k6/latest/get-started/installation/)
* run k6 load test
   ``` shell
    # fix the game-name to your game name
    k6 run ./tests/{game-name}/{game-name}.js
  ```

## Proto file Manage

* install [buf](https://buf.build/docs/installation)

* manage proto file
  ```shell
   #  generate proto file
    buf generate
  ```
  ```shell
   # use buf Schema Registry to manage proto file
   # you need to sign up and login to buf Schema Registry,follow the steps below:
   # https://buf.build/docs/tutorials/getting-started-with-bsr#prerequisites
    buf registry login username 
   # push proto file to buf Schema Registry
    buf push
  ```
* generate SDKS for different languages
    * visit https://buf.build/everyjoy/{game-name}/sdks
    * choose the language you want to generate, and follow the cmd to import the SDKS to your project.
  

      