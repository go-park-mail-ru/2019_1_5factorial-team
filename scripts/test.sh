#!/usr/bin/env bash

check_cover()
{
    go tool cover -func c.out

    go test -coverpkg=./... -coverprofile=c.out ./...
}

./start_dbs_test.sh

check_cover