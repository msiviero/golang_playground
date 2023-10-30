package main

import (
	"dev.msiviero/example/internal/app"
)

func main() {
	app, err := InitializeApp()

	if err != nil {
		panic(err)
	}

	app.Run()
}

func InitializeApp() (app.App, error) {
	return app.App{}, nil
}
