#!/bin/sh

set -e

docker-compose up &

docker_pid=$!

trap "docker-compose down" SIGINT SIGABRT

while docker-compose exec ipfs cat /data/ipfs/api; [[ $? -eq 1 ]]; do sleep 1; done

docker-compose exec ipfs ipfs config --json API.HTTPHeaders.Access-Control-Allow-Origin "[\"*\"]"

docker-compose exec ipfs ipfs config --json API.HTTPHeaders.Access-Control-Allow-Methods "[\"PUT\", \"GET\", \"POST\"]"

docker-compose exec ipfs ipfs config --json API.HTTPHeaders.Access-Control-Allow-Credentials "[\"true\"]"

wait $docker_pid
