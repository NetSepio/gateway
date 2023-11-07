#!/bin/bash
export PATH=$PATH:$(pwd)
mkdir .aptos;
echo $APTOS_CONFIG | base64 -d > .aptos/config.yaml;
./gateway;