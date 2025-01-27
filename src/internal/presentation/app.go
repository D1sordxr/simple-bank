package presentation

type Server interface {
	Run()
}

type App struct {
	Server
}

func NewApp(s Server) *App {
	return &App{Server: s}
}

func (a *App) RunApp() {
	a.Server.Run()
}
