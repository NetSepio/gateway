#!/bin/bash
cd /home/ubuntu/Gateway
docker stop gateway && docker rm gateway
docker image rm 560914023379.dkr.ecr.us-east-2.amazonaws.com/netsepio-gateway --force
docker run -d --name="gateway" \
--add-host="postgres:host-gateway" \
    -v $(pwd):/app/logs \
    --env-file=.env \
    560914023379.dkr.ecr.us-east-2.amazonaws.com/netsepio-gateway