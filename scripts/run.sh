#!/bin/bash
cd /home/ec2-user/plant-system-api
docker-compose -f docker/development/docker-compose.yml build --no-cache
docker-compose -f docker/development/docker-compose.yml up -d