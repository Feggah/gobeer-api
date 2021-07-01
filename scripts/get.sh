#!/bin/bash

curl -v http://localhost:4000/v1/beer/$1 | jq
