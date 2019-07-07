#!/usr/bin/env bash

docker build -t benjaminslabbert/grpc-linkerd-k8s-example-grpc-server -f apps/grpc/server/Dockerfile .
