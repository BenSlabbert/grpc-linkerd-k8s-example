#!/bin/bash

docker build -t benjaminslabbert/grpc-linkerd-k8s-example-grpc-client:$1 -f apps/grpc/client/Dockerfile .
