#!/bin/sh

git submodule update --init --remote

git submodule foreach git pull origin $1
