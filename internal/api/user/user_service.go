package user

import pb "dev.msiviero/example/internal/grpc_gen"

type UserService interface {
	GetUser() pb.UserMessage
}

type UserServiceImpl struct {
}

func NewUserService() UserService {
	return UserServiceImpl{}
}

func (UserServiceImpl) GetUser() pb.UserMessage {
	return pb.UserMessage{Age: 40, Name: "Marco"}
}
