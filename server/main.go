package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/dcmcand/grpc-demo/pb"
	"google.golang.org/grpc"
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

	log.Println("Server Starting")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))

	log.Printf("Server listening on %d", listeningPort)
	if err != nil {
		log.Fatalf("Failed to start listener: %v", err)
	}

	s := Server{}
	grpcServer := grpc.NewServer(grpc.MaxSendMsgSize(1024 * 1024 * 1024))
	pb.RegisterEchoServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}

}

type Server struct {
	pb.UnimplementedEchoServer
}

func (s *Server) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	log.Printf("Recieved Request: %+v", req)
	messages := []string{}
	for i := 0; i < int(req.GetTimes()); i++ {
		messages = append(messages, req.GetMessage())
	}
	res := &pb.EchoResponse{
		ResponseMessage: messages,
	}
	return res, nil
}
