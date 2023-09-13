#!/bin/sh

# Run the application
set -e

echo "run db migration"
/app/migrate -path /app/db/migrations -database "$DB_SOURCE" -verbose up

echo "start"
exec "$@"
