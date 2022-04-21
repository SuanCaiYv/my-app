#!/bin/zsh
docker container prune
docker image rm my-app_backend
docker image rm my-app_frontend
docker compose up