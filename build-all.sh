#!/bin/bash

VERSION=1.0.4

echo "Building all Images"

docker build -t benjaminslabbert/grpc-linkerd-k8s-example-gateway-client:${VERSION} -f apps/gateway/client/Dockerfile .
docker build -t benjaminslabbert/grpc-linkerd-k8s-example-gateway-server:${VERSION} -f apps/gateway/server/Dockerfile .
docker build -t benjaminslabbert/grpc-linkerd-k8s-example-grpc-client:${VERSION} -f apps/grpc/client/Dockerfile .
docker build -t benjaminslabbert/grpc-linkerd-k8s-example-grpc-server:${VERSION} -f apps/grpc/server/Dockerfile .

echo "#########################"
echo "# Pushing to Docker hub #"
echo "#########################"

echo "$DOCKER_HUB_PASSWORD" | docker login -u "$DOCKER_HUB_USERNAME" --password-stdin

docker push benjaminslabbert/grpc-linkerd-k8s-example-gateway-client:${VERSION}
docker push benjaminslabbert/grpc-linkerd-k8s-example-gateway-server:${VERSION}
docker push benjaminslabbert/grpc-linkerd-k8s-example-grpc-client:${VERSION}
docker push benjaminslabbert/grpc-linkerd-k8s-example-grpc-server:${VERSION}
