#!/usr/bin/env bash

docker build -t benjaminslabbert/grpc-linkerd-k8s-example-grpc-client -f apps/grpc/client/Dockerfile .
