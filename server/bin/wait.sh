#!/bin/sh

set -e

cmd="$@"

until PGPASSWORD=chat psql -h "postgres" -U "postgres" -p "5432" -c '\l'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

>&2 echo "Postgres is up - executing command"

exec $cmd
