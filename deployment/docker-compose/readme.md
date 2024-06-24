# 10分钟搭建自动部署(CD)的单机集群环境

### 相关工具

* [docker](https://docs.docker.com/engine/install/)
* [docker-compose](https://docs.docker.com/compose)
* [watchtower](https://github.com/containrrr/watchtower)
* [harbor](https://goharbor.io/docs/2.3.0/install-config/)

### 介绍

本文档介绍了如何在单机部署一个包含私有仓库镜像的`docker-compose`集群环境,并实现服务的全自动更新部署，
你需要了解以上工具的知识点,以及部署自己的私有仓库。

### 适用场景

本地开发测试环境的全自动化部署:
> 代码提交 -> jenkins构建docker镜像 -> jenkins push docker镜像 -> **docker-compose检测到更新后自动更新重启服务**

### 环境配置

1. [安装docker](https://docs.docker.com/engine/install/),注意带上`docker-compose-plugin` 选项(安装教程默认会带)。

2. 修改`daemon.json`文件,[daemon.json在哪里?](https://docs.docker.com/config/daemon/#configuration-file)
   ```json
   {
    "builder": {
        "gc": {
            "defaultKeepStorage": "200GB",
            "enabled": true
        }
    },
   // 如果你的仓库配置的是本地私有证书，需要配置以下内容
    "insecure-registries": [
        "<your_private_registry>//note: without http:// or https://"
    ],
    // 国内需要配置代理，否则无法访问到dockerhub,注意这儿要忽略自己的私有仓库
    "proxies": {
        "http-proxy": "http://<your_proxy>:7890",
        "https-proxy": "http://<your_proxy>:7890",
        "no-proxy": "*.<your_private_registry>"
    }
   }
   ```
3. 重启docker使配置生效
   ```shell
    sudo systemctl daemon-reload
    sudo systemctl restart docker
    ```
4. 登录私有仓库
   ```shell
   // 推荐给当前环境创建一个独立的robot账号密码,具体可以在`harbor`的用户管理中创建
   // 注意<your_private_registry> 不能包含http:// or https://
   docker login <your_private_registry>
   ```

### 配置自己的集群

```yaml
   version: '3'
   services:
     app:
       env_file:
         - .env
       container_name: app
       image: ${REGISTRY}/app:latest
       labels:
         # Enable watchtower for this container
         # 添加这个标签，watchtower会自动更新这个容器
         - "com.centurylinklabs.watchtower.enable=true"
       ports:
         - "8888:8888"
       restart: always
       networks:
         - app-network
     watchtower:
       container_name: watchtower
       image: containrrr/watchtower
       volumes:
         - /root/.docker/config.json:/config.json
         - /var/run/docker.sock:/var/run/docker.sock
       command:
         --interval 30
         --cleanup
         --label-enable
   networks:
     app-network:
```

### 启动集群

   ```shell
    docker-compose up -d
    ## 查看当前watchtower是否正常运行
    docker logs -f watchtower
   ```


