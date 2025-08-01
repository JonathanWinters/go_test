#!/bin/bash

# check for any running containers (they cannot be pruned)
source ./stop-docker.sh
# remove all images
docker system prune -af

# remove all volumes
docker volume prune -f

# remove any dangling volumes
docker volume rm $(docker volume ls -qf dangling=true)
