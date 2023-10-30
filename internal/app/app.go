package app

type App struct {
	Name string
}

func (a *App) GetName() string {
	return a.Name
}
