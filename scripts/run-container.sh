#!/bin/bash

# Run Docker image in a container
docker container rm -f pixelart
docker run --name pixelart -p 8080:8080 --env-file ./backend/.env.list -d -v pixelart-backups:/backups pixelart
