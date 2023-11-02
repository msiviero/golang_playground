package user

import (
	"context"

	pb "dev.msiviero/example/internal/grpc_gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserRoute struct {
	pb.UnimplementedUserRouteServer
	userService UserService
}

func NewUserRoute(userService UserService) UserRoute {
	return UserRoute{userService: userService}
}

func (userRoute *UserRoute) GetUsers(in *pb.EmptyRequest, stream pb.UserRoute_GetUsersServer) error {
	userRoute.userService.GetUsers(func(user User) {
		stream.Send(&pb.UserMessage{
			Id:    user.Id,
			Age:   user.Age,
			Name:  user.Name,
			Email: user.Email,
		})
	})
	return nil
}

func (userRoute *UserRoute) GetUser(ctx context.Context, in *pb.UserRequest) (*pb.UserMessage, error) {
	user, err := userRoute.userService.GetUser(in.Id)

	if err != nil {
		return nil, status.Error(codes.NotFound, "User not found")
	}

	message := userToMessage(&user)
	return &message, nil
}

func (userRoute *UserRoute) PutUser(ctx context.Context, in *pb.UserMessage) (*pb.EmptyReply, error) {
	err := userRoute.userService.PutUser(messageToUser(in))
	return &pb.EmptyReply{}, err
}

func messageToUser(message *pb.UserMessage) User {
	return User{
		Id:    message.Id,
		Age:   message.Age,
		Name:  message.Name,
		Email: message.Email,
	}
}

func userToMessage(user *User) pb.UserMessage {
	return pb.UserMessage{
		Id:    user.Id,
		Age:   user.Age,
		Name:  user.Name,
		Email: user.Email,
	}
}
