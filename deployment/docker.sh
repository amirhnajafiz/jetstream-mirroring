#!/bin/sh

# docker build
docker compose build --no-cache

# docker up
docker compose up -d
