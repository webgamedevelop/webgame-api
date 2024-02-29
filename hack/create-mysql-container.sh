#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

docker pull --platform linux/amd64 mysql:5.7.43

docker run --name mysql -itd -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 \
    -v /usr/bin/qemu-x86_64-static:/usr/bin/qemu-x86_64-static \
    mysql:5.7.43 \
    --character-set-server=utf8mb4 \
    --collation-server=utf8mb4_unicode_ci
