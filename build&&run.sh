#!/bin/bash
pwd
whoami

echo "build and run 5factorial backend server via docker"

if [[ -d "/var/www/media/factorial" ]]
then
    echo "Directory /var/www/media/factorial exists."
else
    echo "Error: Directory /var/www/media/factorial does not exists. Creating dir"
    mkdir -p /var/www/media/factorial
    chmod -R 777 /var/www/media/factorial
fi

if [[ $1 == '-r' ]]
then
    echo "running docker"
    docker run -d -p 5051:5051 -v /var/www/media/factorial:/var/www/media/factorial --name 5factorial-main-back 5factorial-back:latest
elif [[ $1 == '-b' ]]
then
    echo "building docker"
    docker build -t 5factorial-back .
else
    echo "build && run docker"
    docker run -d -p 5051:5051 -v /var/www/media/factorial:/var/www/media/factorial --name 5factorial-main-back 5factorial-back:latest
    docker build -t 5factorial-back .
fi