package main

import (
	"context"
	"log"
	"net/http"

	"connectrpc.com/connect"
	greetv1 "example/connect-go-example/gen/greet/v1"
	"example/connect-go-example/gen/greet/v1/greetv1connect"
)

func main() {
	client := greetv1connect.NewGreetServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
		connect.WithGRPC(), // use gRPC protocol instead of Connect protocol
	)
	res, err := client.Greet(
		context.Background(),
		connect.NewRequest(&greetv1.GreetRequest{Name: "Jane"}),
	)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(res.Msg.Greeting)
}
