// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"dev.msiviero/example/internal/api/user"
	"dev.msiviero/example/internal/app"
	"dev.msiviero/example/internal/grpc_server"
)

// Injectors from wire.go:

func InitializeApp() app.App {
	userRoute := user.NewUserRoute()
	grpcServer := grpc_server.NewGrpcServer(userRoute)
	appApp := app.NewApp(grpcServer)
	return appApp
}
