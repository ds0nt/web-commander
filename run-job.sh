#!/bin/bash

cd "$(dirname $0)/jobs"

cat "$1.js" | docker run -i -v `pwd`:/root node