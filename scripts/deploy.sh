#!/bin/bash

docker login -u $DOCKER_USER -p $DOCKER_PASS
docker push maartje/myip.ninja