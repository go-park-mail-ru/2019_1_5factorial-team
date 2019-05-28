#!/usr/bin/env bash

mongoClient="mongo --port 27051"

$mongoClient << EOF
use user
db.createCollection("profile")
db.profile.createIndex({"email": 1}, {unique: true})
db.profile.createIndex({"nickname": 1}, {unique: true})
EOF


ssh -i 5factorial_id_rsa.pem ubuntu@95.163.180.8 << EOF
cd backend/2019_1_5factorial-team
git fetch --all
git pull
EOF