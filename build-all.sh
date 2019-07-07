#!/bin/bash

echo "Building all Images"

echo "###########################"
echo "# Building Gateway Client #"
echo "###########################"
./build-gateway-client.sh

echo "###########################"
echo "# Building Gateway Server #"
echo "###########################"
./build-gateway-server.sh

echo "########################"
echo "# Building gRPC Client #"
echo "########################"
./build-grpc-client.sh

echo "########################"
echo "# Building gRPC Server #"
echo "########################"
./build-grpc-server.sh
