#!/bin/bash

curl -v -X "DELETE" "http://localhost:4000/v1/beer/$1" \
    -H 'Accept: application/json' \
    -H 'Content-Type: application/json'
