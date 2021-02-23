#!/bin/bash

container_id=$(docker ps -q -f name=ur-last-fm-postgres)

docker exec --user postgres --interactive --tty $container_id psql lastfm
