package user

import (
	"context"

	pb "dev.msiviero/example/internal/grpc"
)

type UserRoute struct {
	pb.UnimplementedUserRouteServer
	userService UserService
}

func NewUserRoute(userService UserService) UserRoute {
	return UserRoute{userService: userService}
}

func (userRoute *UserRoute) GetUser(context context.Context, in *pb.EmptyRequest) (*pb.UserMessage, error) {
	out := userRoute.userService.GetUser()
	return &out, nil
}
