#!/bin/bash

set -e -u -x

find demo-app/ -name "Dockerfile*" \
	-exec echo "Starting linting " {} \;
	-exec dockerfile_lint -f {} \;

echo "Linted all files with success"

 
