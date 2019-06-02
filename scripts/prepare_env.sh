#!/usr/bin/env bash

check_dir() {
if [[ -d "$1" ]]
then
    echo "-- Directory $1 exists"
else
    echo "-- Directory $1 does not exists. Creating dir..."
    mkdir -p $1
    chmod -R 777 $1
fi
return 0
}

copy_file_from_template() {
cp $1.template $2

return 0
}

copy_configs() {
echo "-- $1/auth_grpc_config -> $2"
copy_file_from_template $1/auth_grpc_config.json $2/auth_grpc_config.json

echo "-- $1/chat_config -> $2"
copy_file_from_template $1/chat_config.json $2/chat_config.json

echo "-- $1/cookie_config -> $2"
copy_file_from_template $1/cookie_config.json $2/cookie_config.json

echo "-- $1/cors_config -> $2"
copy_file_from_template $1/cors_config.json $2/cors_config.json

echo "-- $1/db_config -> $2"
copy_file_from_template $1/db_config.json $2/db_config.json

echo "-- $1/game_config -> $2"
copy_file_from_template $1/game_config.json $2/game_config.json

echo "-- $1/logrus_config -> $2"
copy_file_from_template $1/logrus_config.json $2/logrus_config.json

echo "-- $1/static_server_config -> $2"
copy_file_from_template $1/static_server_config.json $2/static_server_config.json

echo "-- $1/user_faker_config -> $2"
copy_file_from_template $1/user_faker_config.json $2/user_faker_config.json

echo "-- $1/donation_server_config -> $2"
copy_file_from_template $1/donation_server_config.json $2/donation_server_config.json

return 0
}

# /////////////////////////////////////////
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
check_dir "/var/www/media/factorial"

echo "== OK"

# /////////////////////////////////////////
echo "Creating dir for configs..."

# проверка папки для хранения конфигов
check_dir "/etc/5factorial/auth"
check_dir "/etc/5factorial/chat"
check_dir "/etc/5factorial/game"
check_dir "/etc/5factorial/core"
check_dir "/etc/5factorial/donation"

echo "== OK"

# /////////////////////////////////////////
echo "Copying config files from templates..."

echo "/etc/5factorial/auth"
copy_configs ./../configs/auth /etc/5factorial/auth
echo "-- OK"

echo "/etc/5factorial/chat"
copy_configs ./../configs/chat /etc/5factorial/chat
echo "-- OK"

echo "/etc/5factorial/game"
copy_configs ./../configs/game /etc/5factorial/game
echo "-- OK"

echo "/etc/5factorial/core"
copy_configs ./../configs/core /etc/5factorial/core
echo "-- OK"

echo "/etc/5factorial/donation"
copy_configs ./../configs/donation /etc/5factorial/donation

echo "/etc/5factorial"
copy_configs ./../configs /etc/5factorial
echo "== OK"

# /////////////////////////////////////////
echo "Creating dir for db data..."

# проверка папки для хранения данных БД
check_dir "/data/db"
check_dir "/data/db/user"
check_dir "/data/db/session"
check_dir "/data/db/chat_global"

echo "== OK"

# /////////////////////////////////////////
echo "Creating docker network"

sudo docker network create 5factorial-net