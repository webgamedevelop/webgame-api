#!/usr/bin/env bash

docker rm -f mysql
docker run -itd --name mysql \
    -e MYSQL_ROOT_PASSWORD=123456 \
    -p 3306:3306 \
    mysql:5.7.44 \
    --character-set-server=utf8mb4 \
    --collation-server=utf8mb4_unicode_ci

docker logs -f mysql
