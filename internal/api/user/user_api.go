package user

import (
	"context"

	pb "dev.msiviero/example/internal/grpc_gen"
)

type UserRoute struct {
	pb.UnimplementedUserRouteServer
	userService UserService
}

func NewUserRoute(userService UserService) UserRoute {
	return UserRoute{userService: userService}
}

func (userRoute *UserRoute) GetUser(ctx context.Context, in *pb.UserRequest) (*pb.UserMessage, error) {
	out := userRoute.userService.GetUser(in.Id)
	return &out, nil
}

func (userRoute *UserRoute) PutUser(ctx context.Context, in *pb.UserMessage) (*pb.EmptyReply, error) {
	err := userRoute.userService.PutUser(in)
	return &pb.EmptyReply{}, err
}
