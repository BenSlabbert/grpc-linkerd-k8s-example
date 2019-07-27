package main

import (
	"context" // Use "golang.org/x/net/context" for Golang version <= 1.6
	"flag"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc-linkerd-k8s-example/grpc/server"
	"grpc-linkerd-k8s-example/pb"
	"net"
	"net/http"
	"os"
	"os/signal"
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	serverPort := os.Getenv("GRPC_PORT")

	if serverPort == "" {
		serverPort = "50051"
	}

	// register gRPC server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", serverPort))

	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}

	// metrics
	go func() {
		log.Info("Setting up metrics")
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":2112", nil)
	}()

	s := grpc.NewServer()
	pb.RegisterYourServiceServer(s, &server.EchoServer{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	go func() {
		log.Infof("Starting gRPC server")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to server: %v", err)
		}
	}()

	// wait for Ctrl + C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, os.Kill)

	// Block until signal is received
	sig := <-ch

	log.Infof("Stopping the server. OS Signal: %v", sig)
	s.Stop()

	log.Infof("Closing the listener")
	_ = lis.Close()

	return nil
}

func main() {
	flag.Parse()

	if err := run(); err != nil {
		log.Fatal(err)
	}
}
