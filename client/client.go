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

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}
	listeningPort, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("could not use listening port: %v", err)
	}

	var conn *grpc.ClientConn
	conn, err = grpc.Dial(fmt.Sprintf(":%d", listeningPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to server: %v", err)
	}
	defer conn.Close()
	c := pb.NewEchoClient(conn)
	log.Printf("Sending \"Hello There\"")
	req := &pb.EchoRequest{
		Message: "Hello There",
		Times:   43,
	}
	response, err := c.Echo(context.Background(), req)
	if err != nil {
		log.Fatalf("Response failed: %v", err)
	}
	for i, msg := range response.GetResponseMessage() {
		fmt.Printf("%d. %s\n", i, msg)
	}
}
