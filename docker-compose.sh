#!/bin/sh

set -e

docker-compose up &

docker_pid=$!

trap "docker-compose down" SIGINT SIGABRT

wait $docker_pid
