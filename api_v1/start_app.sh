# !/bin/bash

until mysqladmin ping -h db --silent; do
  echo 'waiting for mysqld to be connectable...'
  sleep 2
done

echo "app is starting...!"
exec go run main.go