#!/bin/bash

# Script to build the Docker image
docker image rm -f pixelart
docker build --rm -t pixelart .
