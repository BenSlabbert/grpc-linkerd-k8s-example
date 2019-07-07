#!/bin/bash

echo "$DOCKER_HUB_PASSWORD" | docker login -u "$DOCKER_HUB_USERNAME" --password-stdin

docker push benjaminslabbert/grpc-linkerd-k8s-example-gateway-client
docker push benjaminslabbert/grpc-linkerd-k8s-example-gateway-server
docker push benjaminslabbert/grpc-linkerd-k8s-example-grpc-client
docker push benjaminslabbert/grpc-linkerd-k8s-example-grpc-server
