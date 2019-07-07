#!/bin/bash

docker build -t benjaminslabbert/grpc-linkerd-k8s-example-grpc-server:$1 -f apps/grpc/server/Dockerfile .
