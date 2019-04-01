#!/usr/bin/env bash

cd ..
cd third-party

pwd

docker-compose up --build "$@"