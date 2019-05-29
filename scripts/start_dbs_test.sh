#!/usr/bin/env bash

./stop.sh

cd ..
cd third-party

#docker-compose up --build "$@" --file docker-compose_test.yml
docker-compose --file docker-compose_test.yml up --build "$@" -d
