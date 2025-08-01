#!/bin/bash

# passing the SSH key as an environment variable allows
# the dockerfile access to the key without requiring that
# the key be part of the build context. You can leave the
# ssh key in it's regular spot!
mkdir -p .ssh
echo "$(cat ~/.ssh/id_rsa)" > .ssh/id_rsa
echo $'Host *\n StrictHostKeyChecking no\n UserKnownHostsFile=/dev/null' > .ssh/config
chmod 600 .ssh/id_rsa

#default branch to dev for images
export BRANCH=dev

#if no arguments, default to basic dependencies
if [ $# -eq 0 ]; then
    docker compose -f docker/db-docker-compose.yml up
else
    case $1 in
        all)
            docker compose -f docker/db-docker-compose.yml up --build
            ;;
        all-detach)
            docker compose -f -f docker/db-docker-compose.yml up --detach --build
            ;;
        *)
            printf 'No match for "%s"\n' "$1"
            ;;
    esac
fi
