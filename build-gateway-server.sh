#!/bin/bash

docker build -t benjaminslabbert/grpc-linkerd-k8s-example-gateway-server -f apps/gateway/server/Dockerfile .
