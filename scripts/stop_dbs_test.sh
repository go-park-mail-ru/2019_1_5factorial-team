#!/usr/bin/env bash

./stop.sh

cd ..
cd third-party

docker-compose --file docker-compose_test.yml down -v
