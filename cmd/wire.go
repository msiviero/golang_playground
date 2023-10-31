//go:build wireinject
// +build wireinject

package main

import (
	"dev.msiviero/example/internal/api/user"
	"dev.msiviero/example/internal/app"
	"dev.msiviero/example/internal/grpc_server"
	"github.com/google/wire"
)

func InitializeApp() app.App {

	wire.Build(
		app.NewApp,
		grpc_server.NewGrpcServer,
		user.NewUserRoute,
		user.NewUserService,
	)

	return app.App{}
}
