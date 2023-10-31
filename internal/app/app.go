package app

import (
	"dev.msiviero/example/internal/grpc_server"
)

type App struct {
	server grpc_server.GrpcServer
}

func NewApp(server grpc_server.GrpcServer) App {
	return App{server: server}
}

func (app *App) Run() {
	app.server.Start()
}
