package user

import pb "dev.msiviero/example/internal/grpc"

type UserService struct {
}

func NewUserService() UserService {
	return UserService{}
}

func (UserService) GetUser() pb.UserMessage {
	return pb.UserMessage{Age: 40, Name: "Marco"}
}
