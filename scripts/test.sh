#!/usr/bin/env bash

check_cover()
{
    cd ..
    sleep 15
    go test -coverpkg=./internal/... -coverprofile=cover.out.tmp ./internal/...
    cat cover.out.tmp | grep -v -E "_easyjson.go|.pb.go|chat|cmd|HelloWorld" > cover.out
    go tool cover -func cover.out
}

show_coverage()
{
    go tool cover -html=cover.out -o cover.html
    open cover.html
}

./start_dbs_test.sh ; check_cover

cd scripts
./stop_dbs_test.sh
#docker stop 5factorial-back-auth-1            5factorial-user-mongo-1    5factorial-session-mongo-1 \
#            5factorial-chat-global-mongo-test 5factorial-session-mongo-3 5factorial-session-mongo-3 \
#            5factorial-session-mongo-2        5factorial-user-mongo-2    5factorial-user-mongo-3