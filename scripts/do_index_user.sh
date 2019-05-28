#!/usr/bin/env bash

mongoClient="mongo --port 27051"

$mongoClient << EOF
use user
db.createCollection("profile")
db.profile.createIndex({"email": 1}, {unique: true})
db.profile.createIndex({"nickname": 1}, {unique: true})
EOF