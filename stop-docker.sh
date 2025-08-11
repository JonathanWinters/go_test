#!/bin/bash

# check for any running docker processess
docker_ps="$(docker ps -q)"
for ps in ${docker_ps}
do

	# log it
	echo -e "docker stop ${ps}"

	# stop it gracefully
	docker stop "${ps}"
done

# give it a copule seconds
echo -e "chilling for 2 seconds..."
sleep 2

docker_ps="$(docker ps -q)"
for ps in ${docker_ps}
do
	# log it
	echo -e "docker rm --force ${ps}"

	# remove it
	docker rm --force "${docker_ps}"
done

# log any remaining processes
docker ps -a
