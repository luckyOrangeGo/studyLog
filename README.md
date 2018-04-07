# 20180405 && 0406 Learn Docker

## 学习计划

---
学习Docker和其衍生品的使用

## 0405 学习记录

Docker 指令图
![COM of Docker](https://github.com/luckyOrangeGo/studyLog/blob/master/Docker/dockercmd.jpg)

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

#从当前目录的Dockerfile文件中建立IMAGE
docker image build -t $REPOSITORY[:TAG] .

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
>Example of Debian Dockerfile

```Dockerfile
# NOTE: this example is taken from the default Dockerfile for the official nginx Docker Hub Repo
# https://hub.docker.com/_/nginx/
# NOTE: This file is slightly different then the video, because nginx versions have been updated
#       to match the latest standards from docker hub... but it's doing the same thing as the video
#       describes
FROM debian:stretch-slim
# all images must have a FROM
# usually from a minimal Linux distribution like debian or (even better) alpine
# if you truly want to start with an empty container, use FROM scratch

ENV NGINX_VERSION 1.13.6-1~stretch
ENV NJS_VERSION   1.13.6.0.1.14-1~stretch
# optional environment variable that's used in later lines and set as envvar when container is running

RUN apt-get update \
  && apt-get install --no-install-recommends --no-install-suggests -y gnupg1 \
  && \
  NGINX_GPGKEY=573BFD6B3D8FBC641079A6ABABF5BD827BD9BF62; \
  found=''; \
  for server in \
    ha.pool.sks-keyservers.net \
    hkp://keyserver.ubuntu.com:80 \
    hkp://p80.pool.sks-keyservers.net:80 \
    pgp.mit.edu \
  ; do \
    echo "Fetching GPG key $NGINX_GPGKEY from $server"; \
    apt-key adv --keyserver "$server" --keyserver-options timeout=10 --recv-keys "$NGINX_GPGKEY" && found=yes && break; \
  done; \
  test -z "$found" && echo >&2 "error: failed to fetch GPG key $NGINX_GPGKEY" && exit 1; \
  apt-get remove --purge -y gnupg1 && apt-get -y --purge autoremove && rm -rf /var/lib/apt/lists/* \
  && echo "deb http://nginx.org/packages/mainline/debian/ stretch nginx" >> /etc/apt/sources.list \
  && apt-get update \
  && apt-get install --no-install-recommends --no-install-suggests -y \
            nginx=${NGINX_VERSION} \
            nginx-module-xslt=${NGINX_VERSION} \
            nginx-module-geoip=${NGINX_VERSION} \
            nginx-module-image-filter=${NGINX_VERSION} \
            nginx-module-njs=${NJS_VERSION} \
            gettext-base \
  && rm -rf /var/lib/apt/lists/*
# optional commands to run at shell inside container at build time
# this one adds package repo for nginx from nginx.org and installs it

RUN ln -sf /dev/stdout /var/log/nginx/access.log \
  && ln -sf /dev/stderr /var/log/nginx/error.log
# forward request and error logs to docker log collector

EXPOSE 80 443
# expose these ports on the docker virtual network
# you still need to use -p or -P to open/forward these ports on host

CMD ["nginx", "-g", "daemon off;"]
# required: run this command when container is launched
# only one CMD allowed, so if there are multiple, last one wins
```

---

>Example of Nginx Dockerfile

```Dockerfile
# this same shows how we can extend/change an existing official image from Docker Hub

FROM nginx:latest
# highly recommend you always pin versions for anything beyond dev/learn

WORKDIR /usr/share/nginx/html
# change working directory to root of nginx webhost
# using WORKDIR is prefered to using 'RUN cd /some/path'

COPY index.html index.html

# nginx container 在后台已经启动了
# I don't have to specify EXPOSE or CMD because they're in my FROM
```

```bash
docker container run -p 80:80 --rm nginx
docker image build -t nginx-with-html .
docker container run -p 80:80 --rm nginx-with-html
docker image tag nginx-with-html:latest cxxvcheng/nginx-with-html:latest
```

---

>Example of Alpine Dockerfile

Dockerfile

```Dockerfile
FROM alpine

COPY ./docker-entrypoint.sh /

ENTRYPOINT ["/docker-entrypoint.sh"]

CMD ["ping", "8.8.8.8"]
```

docker-entrypoint.sh

```bash
#!/bin/sh
set -e

cat /etc/hosts
echo "ok i'm done, now let's ping'"

exec "$@"
```

---

>Build MY OWN Image Dockerfile

[Dockerfile](https://github.com/luckyOrangeGo/studyLog/blob/master/Docker/dockerfile-assignment-1/Dockerfile)

```bash
docker container run --rm -p 80:3000 testnode
docker tag testnode cxxvcheng/testing-node
docker push cxxvcheng/testing-node
docker image rm cxxvcheng/testing-node testnode
docker container run --rm -p 80:3000 cxxvcheng/testing-node
```

[http://localhost/](http://localhost/)

## 0406 学习记录

学习docker Volume 和 docker-compose

---

## Docker Volume

## CMD

---
Usage:  docker volume COMMAND

    create      Create a volume
    inspect     Display detailed information on one or more volumes
    ls          List volumes
    prune       Remove all unused local volumes
    rm          Remove one or more volumes

`docker container run -d --name mysql -e MYSQL_ALLOW_EMPTY_PASSWORD=True -v mysql-db:/var/lib/mysql mysql`

Mapping a host flie or dir to(**and OVERWRITE**) a container file or dir

`docker container run -v /Users/cxxvc/Documents/GitHub:/path/container` (Mac/Linux)

`docker container run -v //c/Users/%USERNAME%/Documents/GitHub:/path/container` **(Windows)**

`$(pwd)` 当前目录(Mac/Linux)

`%cd%` 当前目录**(Windows)**

### Example Jekyll-serve

```bash
docker run -p 80:4000 -v //c/Users/cxxvc/Documents/GitHub/studyLog/Docker/bindmount-sample-1:/site --name jekyll-serve bretfisher/jekyll-serve
```

## Docker Compose learn

[Docker-Compose官方文档](https://docs.docker.com/compose/compose-file/)

[YAML Reference](http://yaml.org/refcard.html)

## docker-compose CLI

---
最常用的两个命令

`docker-compose up` # 设置 volumes/networks 并启动所有容器

`docker-compose down -v` # 停止所有容器并移除 container/volume/network

### Example of docker-compose.yml

---
[durpal image doc](https://hub.docker.com/_/drupal/)

[postgres image doc](https://hub.docker.com/_/postgres/)

docker-compose.yaml

```yaml
version: '2'

services:
  drupal:
    image: drupal
    ports:
      - "8080:80"
    volumes:
      - drupal-modules:/var/www/html/modules
      - drupal-profiles:/var/www/html/profiles
      - drupal-sites:/var/www/html/sites
      - drupal-themes:/var/www/html/themes
  postgres:
    image: postgres
    environment:
      - POSTGRES_PASSWORD=mypasswd

volumes:
  drupal-modules:
  drupal-profiles:
  drupal-sites:
  drupal-themes:
```

CMD

```bash
docker-compose up
docker-compose down -v
```

### docker-compose build

---

在本文件夹中有四个文件

- nginx.Dockerfile
- nginx.conf
- docker-compose.yml
- html

示例网站来源于 [startbootstrap.com](https://startbootstrap.com/template-overviews/agency/)

nginx.Dockerfile

```Dockerfile
FROM nginx:1.13

COPY nginx.conf /etc/nginx/conf.d/default.conf
```

nginx.conf

```conf
server {

  listen 80;

  location / {

    proxy_pass         http://web;
    proxy_redirect     off;
    proxy_set_header   Host $host;
    proxy_set_header   X-Real-IP $remote_addr;
    proxy_set_header   X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header   X-Forwarded-Host $server_name;

  }
}
```

docker-compose.yml

```yaml
version: '2'

# 仅使用 nginx.conf 到 image

services:
  proxy:
    build:
      context: .
      dockerfile: nginx.Dockerfile
    ports:
      - '80:80'
  web:
    image: httpd
    volumes:
      - ./html:/usr/local/apache2/htdocs/
```

`docker-compose up -d` # 设置 volumes/networks 并在后台启动所有容器
`docker-compose down --rmi local` # 停止所有容器并移除 container/volume/network/images

## Docker Swarm learn

---
问题提出

- 我们如何实现容器生命周期的自动化？
- 我们怎样才能轻松地扩展out/in/up/down？
- 我们如何确保我们的容器在失败时重新创建？
- 我们如何在没有停机的情况下更换容器(蓝色/绿色部署)？
- 我们如何控制/追踪容器的起始位置？
- 我们如何创建跨节点虚拟网络？
- 我们如何才能确保只有受信任的服务器运行我们的容器？
- 我们如何存储秘密，密钥和密码并将它们送到正确的容器（并且只有那个容器）？

## Raft

---
>Understandable Distributed Consensus

可理解的分布式共识 [Raft](http://thesecretlivesofdata.com/raft/)

那么什么是分布式共识？我们从一个例子开始......

### 分布式共识

假设我们有一个单节点系统，对于这个例子，你可以将我们的**节点**看作是存储单个值的数据库服务器。我们也有一个**客户端**可以向服务器发送一个值。就一个节点而言，就该价值达成一致或达成共识很容易。但是，如果我们有多个节点，我们如何达成共识？这就是*分布式共识*的问题。

***Raft***是实施分布式共识的协议。我们来看一下它的工作原理。

### 领导者选举

---
一个节点可以处于三种状态之一：

- The Follower state 追随者状态
- The Candidate state 候选人状态
- The Leader state 领导者状态

我们所有的节点都是从追随者状态开始的。如果追随者没有听到领导者的消息，他们可以成为候选人。候选人然后请求其他节点的投票。节点将回复他们的投票。如果候选人从大多数节点获得投票，他将成为领导者。这个过程被称为*领导者选举*。

现在系统的**所有变化都通过领导者**。每个更改都添加为领导者节点日志中的条目。

如果日志条目当前未提交，它不会更新**领导者节点**的值。要提交条目，**领导者节点**首先将其复制到**跟随者节点上**，然后领导等待，直到大多数跟随者节点已经写入该条目。

该条目现在在领导者节点上提交并且节点状态为“5”（大多数）。领导者随后通知追随者该条目已经落实。该集群现在已经就系统状态达成共识。

这个过程称为*日志复制*。

#### 领袖选举

在**Raft**中，有两个控制选举的超时设置。

首先是**选举超时**。

选举超时是一个追随者在成为候选人之前等待的时间。选举超时被随机分配在150ms和300ms之间。选举超时后，追随者成为候选人，并开始新的*选举期限*...为自己投票...并将*请求投票*消息发送给其他节点。如果接收节点尚未在此期间投票，那么它会为候选人投票...并且该节点重置其选举超时。

一旦候选人A获得多数票，它就成为领导者。领导者A开始发送*Append Entrie 追加条目*消息给其追随者B,C。

这些消息以**心跳超时**指定的间隔发送。追随者B,C然后回应每个*Append Entries 追加条目*消息。

这个选举任期将持续到跟随者B,C停止接受心跳并成为候选人。

我们停止领导人节点A，就会开始*重新选举*。

要求得票的多数保证每个任期只能选一个领导人。如果两个节点同时成为候选人，则可能发生分裂投票。让我们来看看一个分割投票的例子...

#### 分割投票

两个节点同时开始选举......并且每个在另一个之前到达单个跟随者节点。现在每个候选人有2票，并且不能再收到这个任期。节点C在第5期获得多数票，因此成为领导者。

### 日志复制

---
一旦选出了领导者，我们需要将对我们系统的所有更改复制到所有节点。这是通过使用被用于心跳追加条目消息来完成。

首先，客户端向领导者发送变更值为“5”。这个改变被附加到领导者的日志中，然后在下一次心跳中将变化发送给追随者。一旦大多数追随者承认它，领导者就会转换到下一个条目，并将响应发送给客户端。

面对网络分区，Raft甚至可以保持一致。

让我们添加一个分区，将A和B从C，D和E中分离出来。由于我们的分割，我们现在有两个不同的领导者。

A, B(领导者) || C(领导者), D, E

让我们添加另一个客户端，并尝试更新两位领导者。

一个客户端会尝试将节点B的值设置为“3”。领导者节点B不能复制到大多数，所以它的日志条目保持未提交。另一个客户端将尝试将领导者节点C的值设置为“8”。这将成功，因为它可以复制到大多数。

现在我们来修复网络分区。

节点B(期限1)将看到较高的选举期限-节点C(期限2)，并下台。节点A和B都将回滚未提交的条目并匹配新领导的日志。我们的日志现在在整个集群中保持一致。

结束
