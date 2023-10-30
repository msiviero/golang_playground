package grpc_server

import (
	"context"
	"log"

	pb "dev.msiviero/example/internal/grpc"
)

type userRouteServer struct {
	pb.UnimplementedUserRouteServer
}

func (UserRouteServer *userRouteServer) GetUser(context context.Context, in *pb.UserRouteRequest) (*pb.UserRouteReply, error) {
	log.Printf("Received: %v", in)
	return &pb.UserRouteReply{User: &pb.UserMessage{Name: "Marco", Age: 40}}, nil
}
