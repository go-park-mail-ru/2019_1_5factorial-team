#!/usr/bin/env bash

check_cover()
{
#    sleep 30
    go test -coverpkg=./... -coverprofile=c.out ./...
    go tool cover -func c.out
}

./start_dbs_test.sh



