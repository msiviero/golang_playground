package grpc_server

import (
	"fmt"
	"net"
	"os"
	"strconv"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"dev.msiviero/example/internal/api/user"
	pb "dev.msiviero/example/internal/grpc_gen"
)

type GrpcServer struct {
	userRoute user.UserRoute
}

func NewGrpcServer(userRoute user.UserRoute) GrpcServer {
	return GrpcServer{userRoute: userRoute}
}

func (s *GrpcServer) Start() {

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", getPort()))

	if err != nil {
		zap.L().Fatal(fmt.Sprintf("failed to listen: %v", err))
	}

	server := grpc.NewServer()

	pb.RegisterUserRouteServer(server, &s.userRoute)

	zap.L().Info(fmt.Sprintf("server listening at %v", listener.Addr()))
	if err := server.Serve(listener); err != nil {
		zap.L().Fatal(fmt.Sprintf("failed to serve: %v", err))
	}
}

func getPort() int64 {
	val, ok := os.LookupEnv("PORT")
	if !ok {
		return 50051
	}
	port, err := strconv.ParseInt(val, 10, 0)

	if err != nil {
		zap.L().Fatal("Error during PORT env variable conversion")
	}
	return port
}
