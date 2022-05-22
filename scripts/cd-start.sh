#!/bin/bash
cd /home/ubuntu/Gateway
docker stop gateway
docker image rm 560914023379.dkr.ecr.us-east-2.amazonaws.com/netsepio-gateway --force
docker run -d --rm \
    --name="gateway" \
    -p 3000:3000 \
    --add-host="postgres:host-gateway" \
    -v $(pwd):/app/logs \
    --env-file=.env \
    560914023379.dkr.ecr.us-east-2.amazonaws.com/netsepio-gateway