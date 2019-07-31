#!/usr/bin/env bash
docker run --name mysql --rm -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 -e MYSQL_DATABASE=test mysql --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci