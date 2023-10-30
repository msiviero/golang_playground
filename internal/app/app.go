package app

import (
	"dev.msiviero/example/internal/grpc_server"
)

type App struct{}

func (a *App) Run() {
	grpc_server.StartGrpcServer()
}

func NewApp() (App, error) {
	return App{}, nil
}
