#!/bin/sh

# Run the application
set -e

echo "run db migration"
source /app/.env
/app/migrate -path /app/db/migrations -database "$DB_SOURCE" -verbose up

echo "start"
exec "$@"
