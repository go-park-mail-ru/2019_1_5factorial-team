#!/usr/bin/env bash

cd ..
cd third-party

docker-compose up --build "$@"
