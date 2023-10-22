#!/bin/bash
set -e

IFS=',' read -ra DBS <<< "$USERS"
for i in "${DBS[@]}"; do
export PGUSER=postgres
psql <<- EOSQL
    CREATE USER ${i};
    CREATE DATABASE ${i};
    GRANT ALL PRIVILEGES ON DATABASE ${i} TO $i;
EOSQL
done
