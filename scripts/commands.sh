#!/bin/bash
set -e

echo "Cloning the repo"
git clone https://github.com/atkinsonbg/go-gmux-unit-testing.git .

# echo "Setting up the Postgres user"
# POSTGRES="psql --username ${POSTGRES_USER}"

# echo "Creating database: ${DBNAME}"
# $POSTGRES <<EOSQL
# CREATE DATABASE "${DBNAME}" OWNER ${POSTGRESUSER};
# EOSQL

# echo "Creating schema..."
# psql -d ${DBNAME} -a -U${POSTGRESUSER} -f ./scripts/init.sql

# echo "Running Go Tests"
# go test -v ./... &