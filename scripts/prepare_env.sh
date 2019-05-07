#!/usr/bin/env bash

pwd
whoami

set -e

# Init
FILE="/tmp/out.$$"
GREP="/bin/grep"

# Make sure only root can run our script
if [[ $EUID -ne 0 ]]; then
   echo "This script must be run as root" 1>&2
   exit 1
fi

echo "Preparing env..."

# /////////////////////////////////////////
echo "Creating dir for static..."

# проверка папки для статики
if [[ -d "/var/www/media/factorial" ]]
then
    echo "-- Directory /var/www/media/factorial exists"
else
    echo "-- Directory /var/www/media/factorial does not exists. Creating dir..."
    mkdir -p /var/www/media/factorial
    chmod -R 777 /var/www/media/factorial
fi

echo "== OK"

# /////////////////////////////////////////
echo "Creating dir for configs..."

# проверка папки для хранения данных БД
if [[ -d "/etc/5factorial" ]]
then
    echo "-- Directory /etc/5factorial exists"
else
    echo "-- Directory /var/www/media/factorial does not exists. Creating dir..."
    mkdir -p /etc/5factorial
    chmod -R 777 /etc/5factorial
fi

echo "== OK"

# /////////////////////////////////////////
echo "Copying config files from templates"

echo "-- cookie_config"
cp ./../configs/cookie_config.json.template /etc/5factorial/cookie_config.json

echo "-- cors_config"
cp ./../configs/cors_config.json.template /etc/5factorial/cors_config.json

echo "-- db_config"
cp ./../configs/db_config.json.template /etc/5factorial/db_config.json

echo "-- static_server_config"
cp ./../configs/static_server_config.json.template /etc/5factorial/static_server_config.json

echo "-- user_faker_config"
cp ./../configs/user_faker_config.json.template /etc/5factorial/user_faker_config.json

echo "-- session_worker_config"
cp ./../configs/session_worker_config.json.template /etc/5factorial/session_worker_config.json

echo "-- logrus_config"
cp ./../configs/logrus_config.json.template /etc/5factorial/logrus_config.json

echo "== OK"

# /////////////////////////////////////////
echo "Creating dir for db data..."

# проверка папки для хранения данных БД
if [[ -d "/data/db" ]]
then
    echo "-- Directory /data/db exists"
else
    echo "-- Directory /data/db does not exists. Creating dir..."
    mkdir -p /data/db
    chmod -R 777 /data/db
fi

if [[ -d "/data/db/user" ]]
then
    echo "-- Directory /data/db/user exists"
else
    echo "-- Directory /data/db/user does not exists. Creating dir..."
    mkdir -p /data/db/user
    chmod -R 777 /data/db/user
fi

if [[ -d "/data/db/session" ]]
then
    echo "-- Directory /data/db/session exists"
else
    echo "-- Directory /data/db/session does not exists. Creating dir..."
    mkdir -p /data/db/session
    chmod -R 777 /data/db/session
fi

if [[ -d "/data/db/chat_global" ]]
then
    echo "-- Directory /data/db/chat_global exists"
else
    echo "-- Directory /data/db/chat_global does not exists. Creating dir..."
    mkdir -p /data/db/chat_global
    chmod -R 777 /data/db/chat_global
fi

echo "== OK"
