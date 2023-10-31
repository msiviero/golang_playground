package grpc_server

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"google.golang.org/grpc"

	"dev.msiviero/example/internal/api/user"
	pb "dev.msiviero/example/internal/grpc"
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
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	pb.RegisterUserRouteServer(server, &s.userRoute)

	log.Printf("server listening at %v", listener.Addr())
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func getPort() int64 {
	val, ok := os.LookupEnv("PORT")
	if !ok {
		return 50051
	}
	v, err := strconv.ParseInt(val, 10, 0)

	if err != nil {
		fmt.Println("Error during conversion")
		panic(err)
	}
	return v
}
