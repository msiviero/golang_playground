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

func (userRoute *UserRoute) GetUser(context context.Context, in *pb.UserRequest) (*pb.UserMessage, error) {
	out := userRoute.userService.GetUser()
	return &out, nil
}

func (userRoute *UserRoute) PutUser(context.Context, *pb.UserMessage) (*pb.EmptyReply, error) {
	out := pb.EmptyReply{}
	return &out, nil
}
