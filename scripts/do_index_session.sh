#!/usr/bin/env bash

mongoClient="mongo --port 27052"

$mongoClient << EOF
use user_session
db.createCollection("user_session")
db.user_session.createIndex({"user_id": 1}, {unique: true})
EOF