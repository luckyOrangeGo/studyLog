# Docker Compose learn

[Docker-Compose官方文档](https://docs.docker.com/compose/compose-file/)

[YAML Reference](http://yaml.org/refcard.html)

## docker-compose CLI

---
Two most common commands are

`docker-compose up` # setup volumes/networks and start all containers

`docker-compose down` # stop all containers and remove cont/vol/net

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