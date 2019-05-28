#!/usr/bin/env bash

check_cover()
{
    cd ../internal/pkg
    sleep 15
    go test -coverpkg=./... -coverprofile=c.out ./...
    go tool cover -func c.out
}

show_coverage()
{
    go test -v -coverprofile cover.out ./...
    go tool cover -html=cover.out -o cover.html
    open cover.html
}

./start_dbs_test.sh ; check_cover

docker stop 5factorial-back-auth-1            5factorial-user-mongo-1    5factorial-session-mongo-1 \
            5factorial-chat-global-mongo-test 5factorial-session-mongo-3 5factorial-session-mongo-3 \
            5factorial-session-mongo-2        5factorial-user-mongo-2    5factorial-user-mongo-3