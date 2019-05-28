#!/usr/bin/env bash

cd ..
cd third-party

docker-compose --file docker-compose_test.yml down -v
