package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "dev.msiviero/example/internal/grpc"
)

type server struct {
	pb.UnimplementedUserRouteServer
}

func (UserRouteServer *server) GetUser(context context.Context, in *pb.UserRouteRequest) (*pb.UserRouteReply, error) {
	log.Printf("Received: %v", in)
	return &pb.UserRouteReply{User: &pb.UserMessage{Name: "Marco", Age: 40}}, nil
}

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", flag.Int("port", 50051, "The server port")))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	server := server{}

	pb.RegisterUserRouteServer(s, &server)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
