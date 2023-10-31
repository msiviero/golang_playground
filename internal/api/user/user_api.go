package user

import (
	"context"
	"log"

	pb "dev.msiviero/example/internal/grpc"
)

type UserRoute struct {
	pb.UnimplementedUserRouteServer
}

func NewUserRoute() UserRoute {
	return UserRoute{}
}

func (UserRouteServer *UserRoute) GetUser(context context.Context, in *pb.EmptyRequest) (*pb.UserMessage, error) {
	log.Printf("Received: %v", in)
	return &pb.UserMessage{Name: "Marco", Age: 40}, nil
}
