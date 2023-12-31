// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"dev.msiviero/example/internal/api/user"
	"dev.msiviero/example/internal/app"
	"dev.msiviero/example/internal/grpc_server"
	"dev.msiviero/example/internal/sql"
)

// Injectors from wire.go:

func InitializeApp() app.App {
	db := sql.NewSql()
	userService := user.NewUserService(db)
	userRoute := user.NewUserRoute(userService)
	grpcServer := grpc_server.NewGrpcServer(userRoute)
	appApp := app.NewApp(grpcServer)
	return appApp
}
