# Docker 镜像操作流程

## Prepare

1. 在 `game` 根目录下添加 `cert.crt` 文件
2. 在 `game` 根目录下添加 `.netrc`
   文件 [如何添加](https://kekxv.github.io/2021/01/06/User%20configuration%20script%20file%20.netrc/)

## Login

```bash
docker login <your private registry>    
 ```

## Build

 ``` shell
 # fix <appname> to service name
 # fix <your private registry> to your private registry
docker buildx build -t <your private registry> /game/rumble:latest --build-arg  NETRC_CONTENT="$NETRC_CONTENT" --build-arg APP_NAME=<appname> -f ./build/package/base/Dockerfile . --push
 ```



