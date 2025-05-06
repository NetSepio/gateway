#!/bin/bash
export PATH=$PATH:$(pwd)
mkdir -p .aptos;
rm -rf .aptos/config.yaml
echo $APTOS_CONFIG | base64 -d > .aptos/config.yaml;
./gateway;
