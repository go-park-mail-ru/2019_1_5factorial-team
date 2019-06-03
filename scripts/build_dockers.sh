#!/usr/bin/env bash

docker build -t smetk/5factorial-auth:`git log -1 --pretty=%h` \
             -t smetk/5factorial-auth:latest \
             -f ./third-party/auth/Dockerfile \
             .
docker build -t smetk/5factorial-core:`git log -1 --pretty=%h` \
             -t smetk/5factorial-core:latest \
             -f ./third-party/server/Dockerfile \
             .
docker build -t smetk/5factorial-worker:`git log -1 --pretty=%h` \
             -t smetk/5factorial-worker:latest \
             -f ./third-party/worker/Dockerfile \
             .
docker build -t smetk/5factorial-chat:`git log -1 --pretty=%h` \
             -t smetk/5factorial-chat:latest \
             -f ./third-party/chat/Dockerfile \
             .
docker build -t smetk/5factorial-game:`git log -1 --pretty=%h` \
             -t smetk/5factorial-game:latest \
             -f ./third-party/game/Dockerfile \
             .

docker build -t smetk/5factorial-donation:`git log -1 --pretty=%h` \
             -t smetk/5factorial-donation:latest \
             -f ./third-party/donation/Dockerfile \
             .

docker push smetk/5factorial-auth   ; docker push smetk/5factorial-core ;
docker push smetk/5factorial-worker ; docker push smetk/5factorial-chat ;
docker push smetk/5factorial-game   ; docker push smetk/5factorial-donation