#!/bin/bash

# if $1  is up, build the docker file and start server
if [ "$1" == "up" ]; then
    docker build -t server . && docker run -p 8080:8080 --name awesomeha-backend server
fi