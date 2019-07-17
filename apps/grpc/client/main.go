package main

import (
	"flag"
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"grpc-linkerd-k8s-example/grpc/client"
	"grpc-linkerd-k8s-example/pb"
	"log"
	"os"
	"time"
)

func run() {
	serverHost := os.Getenv("GRPC_HOST")

	log.Printf("Staring gRPC client to server host: %s", serverHost)

	cc, err := grpc.Dial(serverHost, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	// called at program end
	defer client.CloseConnection(cc)

	c := pb.NewYourServiceClient(cc)

	for {
		go client.CallPing(c)
		go client.CallEcho(c)
		go client.CallEchoServerStream(c)
		go client.CallEchoClientStream(c)
		go client.CallBiDiStream(c)

		time.Sleep(5 * time.Second)
	}
}

func main() {
	flag.Parse()
	defer glog.Flush()

	run()
	glog.Fatal("Failed to call grpc server")
}
