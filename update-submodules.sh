#!/bin/sh

git submodule update --init --remote

git submodule foreach git checkout main