#!/usr/bin/env bash

./stop.sh

cd ..
cd third-party

docker-compose up --build "$@"
