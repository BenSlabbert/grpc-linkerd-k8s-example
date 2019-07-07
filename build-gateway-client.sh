#!/usr/bin/env bash

docker build -t benjaminslabbert/grpc-linkerd-k8s-example-gateway-client -f apps/gateway/client/Dockerfile .
