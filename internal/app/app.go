package app

import (
	"dev.msiviero/example/internal/grpc_server"
	"go.uber.org/zap"
)

type App struct {
	server grpc_server.GrpcServer
}

func NewApp(server grpc_server.GrpcServer) App {
	zap.ReplaceGlobals(zap.Must(zap.NewProduction()))

	return App{server: server}
}

func (app *App) Run() {
	app.server.Start()
}
