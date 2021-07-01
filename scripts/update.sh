#!/bin/bash

while getopts i:n:t:s: flag
do
    case "${flag}" in
        i) id=${OPTARG};;
        n) name=${OPTARG};;
        t) type=${OPTARG};;
        s) style=${OPTARG};;
    esac
done

curl -v -X "PUT" "http://localhost:4000/v1/beer/$id" \
    -H 'Accept: application/json' \
    -H 'Content-Type: application/json' \
    --data '{"name":"'$name'", "type":'$type', "style":'$style'}'
