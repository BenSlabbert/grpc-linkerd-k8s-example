package main

import (
	"context" // Use "golang.org/x/net/context" for Golang version <= 1.6
	"flag"
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc-linkerd-k8s-example/grpc/server"
	"grpc-linkerd-k8s-example/pb"
	"log"
	"net"
	"os"
	"os/signal"
)

var (
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:50051", "gRPC server endpoint")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// register gRPC server
	lis, err := net.Listen("tcp", *grpcServerEndpoint)

	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterYourServiceServer(s, &server.EchoServer{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	go func() {
		log.Println("Starting gRPC server")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to server: %v", err)
		}
	}()

	// wait for Ctrl + C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until signal is received
	<-ch

	log.Println("Stopping the server")
	s.Stop()

	log.Println("Closing the listener")
	_ = lis.Close()

	return nil
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
