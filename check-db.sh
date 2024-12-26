#!/bin/sh

until nc -z -v -w30 db 5432
do
  echo "Waiting for database connection...dasdfasdf"
  sleep 1
done

echo "Database is up, starting the app..."
./blogo --cfg=./config.json
