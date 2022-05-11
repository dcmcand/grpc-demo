package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/dcmcand/grpc-demo/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Change the following two variables to change the client request
var msg string = "Pete and Repeat went out in a boat. Pete fell out. Who is left?" // the message to be echoed
var repeats uint64 = 100                                                           // the number of times to echo it

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "9000"
	}
	listeningPort, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("could not use listening port: %v", err)
	}

	var conn *grpc.ClientConn
	conn, err = grpc.Dial(
		fmt.Sprintf(":%d", listeningPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1024*1024*20)),
	)
	if err != nil {
		log.Fatalf("Could not connect to server: %v", err)
	}
	defer conn.Close()
	c := pb.NewEchoClient(conn)
	log.Printf("Sending \"%s\" %d times", msg, repeats)
	req := &pb.EchoRequest{
		Message: msg,
		Times:   repeats,
	}
	response, err := c.Echo(context.Background(), req)
	if err != nil {
		log.Fatalf("Response failed: %v", err)
	}
	for i, msg := range response.GetResponseMessage() {
		fmt.Printf("%d. %s\n", i, msg)
	}
}
