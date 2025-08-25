#!/bin/sh
set -e

echo "Waiting for Postgres..."
until pg_isready -h db -p 5432 -U "${DB_USER}" -d "${DB_NAME}"; do
  echo "Postgres is unavailable - sleeping"
  sleep 1
done
echo "Postgres is ready"

echo "Running migrations..."
./main -migrate

echo "Running seed..."
./main -seed

echo "Starting app..."
exec ./main