package main

import (
	"flag"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"grpc-linkerd-k8s-example/grpc/client"
	"grpc-linkerd-k8s-example/pb"
	"net/http"
	"os"
	"time"
)

func run() {
	serverHost := os.Getenv("GRPC_HOST")

	log.Printf("Staring gRPC client to server host: %s", serverHost)

	// metrics
	go func() {
		log.Info("Setting up metrics")
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":2112", nil)
	}()

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

	run()
	log.Fatal("Failed to call grpc server")
}
