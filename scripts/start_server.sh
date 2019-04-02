#!/usr/bin/env bash

sleep 5

cd ../third-party/server
pwd

docker-compose up --build "$@"
