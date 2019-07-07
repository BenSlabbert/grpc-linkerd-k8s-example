#!/bin/bash

docker build -t benjaminslabbert/grpc-linkerd-k8s-example-gateway-client:$1 -f apps/gateway/client/Dockerfile .
