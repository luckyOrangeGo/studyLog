# Docker Volume

## CMD

Usage:  docker volume COMMAND

    create      Create a volume
    inspect     Display detailed information on one or more volumes
    ls          List volumes
    prune       Remove all unused local volumes
    rm          Remove one or more volumes

`docker container run -d --name mysql -e MYSQL_ALLOW_EMPTY_PASSWORD=True -v mysql-db:/var/lib/mysql mysql`

Mapping a host flie or dir to(**and OVERWRITE**) a container file or dir

`docker container run -v /Users/cxxvc/Documents/GitHub:/path/container` (Mac/Linux)

`docker container run -v //c/Users/cxxvc/Documents/GitHub:/path/container` **(Windows)**

`$(pwd)` 当前目录(Mac/Linux)

`.` 当前目录**(Windows)**

