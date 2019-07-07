#!/bin/bash

VERSION=1.0.1

echo "Building all Images"

echo "###########################"
echo "# Building Gateway Client #"
echo "###########################"
./build-gateway-client.sh ${VERSION}

echo "###########################"
echo "# Building Gateway Server #"
echo "###########################"
./build-gateway-server.sh ${VERSION}

echo "########################"
echo "# Building gRPC Client #"
echo "########################"
./build-grpc-client.sh ${VERSION}

echo "########################"
echo "# Building gRPC Server #"
echo "########################"
./build-grpc-server.sh ${VERSION}

./docker-push.sh  ${VERSION}
