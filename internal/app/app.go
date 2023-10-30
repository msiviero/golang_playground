package app

import (
	"dev.msiviero/example/internal/grpc_server"
)

type App struct{}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() {
	grpc_server.StartGrpcServer()
}
