#!/bin/bash

VERSION=1.0.2

echo "Building all Images"

echo "###########################"
echo "# Building Gateway Client #"
echo "###########################"
docker build -t benjaminslabbert/grpc-linkerd-k8s-example-gateway-client:${VERSION} -f apps/gateway/client/Dockerfile .


echo "###########################"
echo "# Building Gateway Server #"
echo "###########################"
docker build -t benjaminslabbert/grpc-linkerd-k8s-example-gateway-server:${VERSION} -f apps/gateway/server/Dockerfile .

echo "########################"
echo "# Building gRPC Client #"
echo "########################"
docker build -t benjaminslabbert/grpc-linkerd-k8s-example-grpc-client:${VERSION} -f apps/grpc/client/Dockerfile .

echo "########################"
echo "# Building gRPC Server #"
echo "########################"
docker build -t benjaminslabbert/grpc-linkerd-k8s-example-grpc-server:${VERSION} -f apps/grpc/server/Dockerfile .

echo "#########################"
echo "# Pushing to Docker hub #"
echo "#########################"

echo "$DOCKER_HUB_PASSWORD" | docker login -u "$DOCKER_HUB_USERNAME" --password-stdin

docker push benjaminslabbert/grpc-linkerd-k8s-example-gateway-client:$1
docker push benjaminslabbert/grpc-linkerd-k8s-example-gateway-server:$1
docker push benjaminslabbert/grpc-linkerd-k8s-example-grpc-client:$1
docker push benjaminslabbert/grpc-linkerd-k8s-example-grpc-server:$1
