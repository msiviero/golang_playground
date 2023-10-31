//go:build wireinject
// +build wireinject

package main

import (
	"dev.msiviero/example/internal/app"
	"dev.msiviero/example/internal/grpc_server"
	"github.com/google/wire"
)

func InitializeApp() app.App {

	wire.Build(
		app.NewApp,
		grpc_server.NewUserRouteServer,
		grpc_server.NewGrpcServer,
	)

	return app.App{}
}
