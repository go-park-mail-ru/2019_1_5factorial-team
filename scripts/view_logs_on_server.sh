#!/usr/bin/env bash

# how_to
# ./view_logs_on_server.sh -f -db => -f means do not follow, -db means watch all logs for db
# ./view_logs_on_server.sh -t -all => -t meant follow logs, -all watch all logs on server

server_client="ssh root@78.155.207.69"

if [[ $1 == '-t' ]]
then
    logs="-f"
else
    logs=""
fi


if [[ $2 == "-all" ]]
then
    service=""
elif [[ $2 == "-go" ]]
then
    service="server-go"
elif [[ $2 == "-db" ]]
then
    service="mongo-user mongo-session"
fi

$server_client << EOF
    cd server/2019_1_5factorial-team/third-party/
    docker-compose logs ${logs} ${service}
EOF