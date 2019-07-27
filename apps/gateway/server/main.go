package main

import (
	"context" // Use "golang.org/x/net/context" for Golang version <= 1.6
	"flag"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"grpc-linkerd-k8s-example/grpc/client"
	"grpc-linkerd-k8s-example/pb"
	"net/http"
)

var (
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:50051", "gRPC server endpoint")
	proxyPort          = flag.String("proxy-port", "8081", "Port for proxy to accept requests")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// metrics
	go func() {
		log.Info("Setting up metrics")
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":2112", nil)
	}()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible

	cc, err := grpc.Dial(*grpcServerEndpoint, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	// called at program end
	defer client.CloseConnection(cc)

	c := pb.NewYourServiceClient(cc)

	message, err := client.CallPing(c)

	if err != nil {
		log.Fatalf("Failed to ping grpc server: %v", err)
	}

	log.Printf("Found grpc server at: %v", message)

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err = pb.RegisterYourServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)

	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":"+*proxyPort, serveSwagger(mux))
}

func serveSwagger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/swagger" {
			w.Header().Set("Content-Type", "application/json")
			http.ServeFile(w, r, "/home/ben/go/src/grpc-gateway-example/pb/service.swagger.json")
			return
		}
		h.ServeHTTP(w, r)
	})
}

func main() {
	flag.Parse()

	if err := run(); err != nil {
		log.Fatal(err)
	}
}
