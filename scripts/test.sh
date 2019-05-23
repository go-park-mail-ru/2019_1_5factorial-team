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
./scripts/stop.sh



