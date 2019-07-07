package server

import (
	"context"
	"grpc-linkerd-k8s-example/pb"
	"io"
	"log"
	"net"
	"time"
)

type EchoServer struct {
}

func (EchoServer) Ping(ctx context.Context, req *pb.PingMessage) (*pb.PongMessage, error) {
	log.Println("Ping called")

	ifaces, err := net.Interfaces()

	if err != nil {
		return nil, err
	}

	// handle err
	var ip net.IP
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			return nil, err
		}
		// handle err
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
		}
	}

	if ip == nil {
		return &pb.PongMessage{Ip: "Unknown"}, nil
	}

	return &pb.PongMessage{Ip: ip.String()}, nil
}

func (EchoServer) Echo(ctx context.Context, req *pb.StringMessage) (*pb.StringMessage, error) {
	log.Println("Echo called")

	value := req.GetValue()
	log.Printf("Echo called with: %v", value)

	return &pb.StringMessage{
		Value: value,
	}, nil
}

func (EchoServer) EchoServerStream(req *pb.StringMessage, stream pb.YourService_EchoServerStreamServer) error {
	log.Println("EchoServerStream called")

	value := req.GetValue()

	for i := 0; i < 5; i++ {
		send := stream.Send(&pb.StringMessage{
			Value: value,
		})

		log.Printf("Send result: %v", send)
		time.Sleep(1000 * time.Millisecond)
	}

	return nil
}

func (EchoServer) EchoClientStream(stream pb.YourService_EchoClientStreamServer) error {
	log.Println("EchoClientStream called")

	var lastClientMessage string

	for {
		req, err := stream.Recv()

		if req != nil {
			log.Printf("Message from client: %v", req)
			lastClientMessage = req.GetValue()
		}

		if err == io.EOF {
			// we can respond whenever we want
			log.Println("Finished reading from client")
			return stream.SendAndClose(&pb.StringMessage{Value: lastClientMessage})
		}

		if err != nil {
			log.Fatalf("Error reading from client stream: %v", err)
		}
	}
}

func (EchoServer) EchoBiDiStream(stream pb.YourService_EchoBiDiStreamServer) error {
	log.Println("EchoBiDiStream called")

	//isLastMessage := false
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			//isLastMessage = true
			return nil
		}

		if err != nil && err != io.EOF {
			log.Fatalf("Error while reading client stream")
			return err
		}

		value := req.GetValue()
		log.Printf("Value from client: %v", value)

		sendErr := stream.Send(&pb.StringMessage{Value: value})

		if sendErr != nil {
			log.Fatalf("Error while writing to client")
			return sendErr
		}
	}

	return nil
}
