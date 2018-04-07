# 20180405Learn

## 学习计划

---
学习Docker 的使用

## 学习记录

### 指令

---

`Ctrl + P` 和 `Ctrl + Q` 退出容器而不关闭
exit 退出并关闭容器

ubuntu镜像源
>sudo sed -i 's/archive.ubuntu.com/mirrors.ustc.edu.cn/g' /etc/apt/sources.list

其他镜像源
[mirrors.ustc.edu.cn](http://mirrors.ustc.edu.cn)
[mirrors.163.com](http://mirrors.163.com/)

```bash
#将container保存为image
docker commit $CONTAINER_ID afire/ubuntu_init:v1

docker attach $CONTAINER_ID #连接一个已存在的docker容器
docker stop $CONTAINER_ID #停止docker容器
docker start $CONTAINER_ID #启动docker容器
docker restart $CONTAINER_ID #重启docker容器
docker kill $CONTAINER_ID #强制关闭docker容器
docker logs $CONTAINER_ID #查看docker容器运行日志，确保正常运行
docker inspect $CONTAINER_ID #查看container的属性，比如ip等等
docker rm $CONTAINER_ID      # 删除容器
docker rmi $CONTAINER_ID      # 删除镜像

# $CONTAINER_ID 可用前3位

#show me all running processes
ps aux

winpty #使用windows系统的bash 时，加在前面

#'docker container run' always starts a *new* container
docker container run

#'docker container start' to start an existing stopped one
docker container start
-ai # attach and interactive

#show logs for a specific container
docker container logs

#强制结束***container
docker container rm -f $CONTAINER_ID

#Example Of Changing The Defaults
docker container run -p 80:80 --name webhost -d nginx:1.11 nginx -T
docker container port webhost

#list running processes in specific container
docker container top

# start new container interactively
docker container run -it $IMAGE_NAME bash
# -t simulates a real terminal, like what SSH does
# -i Keep session oopen to receive terminal input
# if run with -it ,it will give you a terminal inside the running container

# ubuntu default CMD is bash,所以不用特别在后面加 bash

# run addition command in existing container
docker container exec -it $CONTAINER_ID bash

# alpine
docker container run -it alpine sh

```

### 学习管理多个Container

---

```bash
docker container run -d -p 3306:3306 --name db -e MYSQL_RANDOM_ROOT_PASSWORD=true mysql

docker container logs db

-->GENERATED ROOT PASSWORD: ***

docker container run -d --name webserber -p 8080:80 httpd

docker container run -d --name proxy -p 80:80 nginx

#docker inspect show metadata about the container(startup config, volumes, networking, etc)
docker container inspect $CONTAINER_ID
--format '{{ .NetworkSettings.IPAddress}}' # A common option for formatting the output of commands using "Go templates"

#docker stats show live performance data for all containers
docker container stats
```

### docker 的网络结构

---

```bash
docker network ls

--network bridge
#默认与HOST IP连接的Docker虚拟网络

--network host
#It gains performance by skipping virtual networks but sacrifices security of container model

–-network none
#removes eth0 and only leaves you with localhost interface in container

docker network create
#Swawns a new virtual network for you to attach containers to

--network driver
#Built-in or 3rd party extensions that give you virtual net work features

#新建网桥连接到新的/已有的容器
docker network create my_app_net

#新的容器
docker container run -d --name new_nginx --network my_app_net nginx

#已有的容器
docker network ls
docker ps
docker network connect $NETWORK_ID $CONTAINER_ID
docker container inspect $CONTAINER_ID

# 断开连接
docker network disconnect $NETWORK_ID $CONTAINER_ID
docker container inspect $CONTAINER_ID
```

#### Docker DNS

由于IP在不同环境经常改变，并不可靠，因此需要DNS server

---
DNS Default Names -- hostname

```bash
#新建两个容器在同一个Network中
docker container run -d --name new_nginx --network my_app_net nginx
docker container run -d --name my_nginx --network my_app_net nginx

# 互相PING 可能需要安装iputils-ping
# apt-get update && apt-get install -y iputils-ping
winpty docker container exec -it new_nginx ping my_nginx
winpty docker container exec -it my_nginx ping new_nginx
```

### Docker Image

---

```bash
docker image ls

#Show layers of changes made in image
docker image history $IMAGE_ID || $REPOSITORY:TAG

# returns JSON metadata about the image
docker image inspect $IMAGE_ID || $REPOSITORY[:TAG]

docker image tag $SOURCE_IMAGE[:TAG] $TARGET_IMAGE[:TAG]

#uploads changed layers to a image registry (default is Hub)
docker image push $REPOSITORY
```

```bash
docker login <server>
# default to logging in Hub

docker logout
```

### Dockerfile

---
