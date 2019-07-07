#!/bin/bash

echo "$DOCKER_HUB_PASSWORD" | docker login -u "$DOCKER_HUB_USERNAME" --password-stdin

docker push benjaminslabbert/grpc-linkerd-k8s-example-gateway-client:$1
docker push benjaminslabbert/grpc-linkerd-k8s-example-gateway-server:$1
docker push benjaminslabbert/grpc-linkerd-k8s-example-grpc-client:$1
docker push benjaminslabbert/grpc-linkerd-k8s-example-grpc-server:$1
