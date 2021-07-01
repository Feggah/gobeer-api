#!/bin/bash

while getopts i:n:t:s: flag
do
    case "${flag}" in
        n) name=${OPTARG};;
        t) type=${OPTARG};;
        s) style=${OPTARG};;
    esac
done

curl -v -POST "http://localhost:4000/v1/beer" \
    -H 'Accept: application/json' \
    -H 'Content-Type: application/json' \
    --data '{"name":"'$name'", "type":'$type', "style":'$style'}'
