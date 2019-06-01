#!/usr/bin/env bash

./stop_dbs_test.sh

cd ..
cd third-party

#docker-compose up --build "$@" --file docker-compose_test.yml
docker-compose --file docker-compose_test.yml up --build "$@" -d
