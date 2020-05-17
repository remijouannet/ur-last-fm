#!/bin/bash

container_id=$(docker run --name ur-last-fm-postgres -p 8081:5432 -e POSTGRES_PASSWORD=password -d postgres)

initdb=false
while ! $initdb
do
    docker exec -u postgres $container_id psql -c "SELECT 1;"
    if [ $? -eq 0 ]
    then
        initdb=true
    fi
    sleep 1
done

docker exec -u postgres $container_id psql -c "CREATE DATABASE lastfm;"
docker exec -u postgres $container_id psql -c "CREATE USER lastfm WITH PASSWORD 'password'"
docker exec -u postgres $container_id psql -c "GRANT ALL PRIVILEGES ON DATABASE lastfm TO lastfm;"
