package main

import (
	"flag"
	"github.com/golang/glog"
	"gopkg.in/resty.v1"
	"log"
	"time"
)

type Message struct {
	Value string `json:"value"`
}

var (
	gateway = flag.String("gateway-server-host", "localhost:8081", "Port for proxy to accept requests")
)

func run() {
	for {
		// GET request
		resp, err := resty.R().Get("http://" + *gateway + "/v1/example/ping")
		// explore response object
		log.Printf("Error: %v", err)
		log.Printf("Response Status Code: %v", resp.StatusCode())
		log.Printf("Response Status: %v", resp.Status())
		log.Printf("Response Time: %v", resp.Time())
		log.Printf("Response Received At: %v", resp.ReceivedAt())
		log.Printf("Response Body: %v", resp) // or resp.String() or string(resp.Body())

		// POST JSON string
		// No need to set content type, if you have client level setting
		resp, err = resty.R().
			SetHeader("Content-Type", "application/json").
			SetBody(`{"value":"testuser"}`).
			SetResult(&Message{}). // or SetResult(AuthSuccess{}).
			Post("http://" + *gateway + "/v1/example/echo")
		log.Printf("Error: %v", err)
		log.Printf("Response Status Code: %v", resp.StatusCode())
		log.Printf("Response Status: %v", resp.Status())
		log.Printf("Response Time: %v", resp.Time())
		log.Printf("Response Received At: %v", resp.ReceivedAt())
		log.Printf("Response Body: %v", resp) // or resp.String() or string(resp.Body())

		time.Sleep(5 * time.Millisecond)
	}
}

func main() {
	flag.Parse()
	defer glog.Flush()

	run()
	glog.Fatal("Failed to call grpc gateway")
}
