#!/bin/sh

TAG="tkitsunai/commango"

docker build --no-cache --rm -t "${TAG}" .
