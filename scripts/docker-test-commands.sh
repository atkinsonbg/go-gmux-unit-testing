#!/bin/bash
set -e

echo "Wait for Postgres to start"
until pg_isready -h localhost -p 5432 -U postgres
do
  echo "Waiting for postgres at: $pg_uri"
  sleep 2;
done
echo "Postgres has started"

echo "Setting up the Postgres user"
POSTGRES="psql --username ${POSTGRESUSER}"

echo "Creating database: ${DBNAME}"
$POSTGRES <<EOSQL
CREATE DATABASE "${DBNAME}" OWNER ${POSTGRESUSER};
EOSQL

echo "Creating schema..."
psql -d ${DBNAME} -a -U${POSTGRESUSER} -f ./scripts/init.sql &
SQL_PID=$!
wait $SQL_PID

echo "Finished with creatin schema"

echo "Running Go Tests"
go test -v ./... &
GO_PID=$!
wait $GO_PID
echo "Finished with go tests"