package presentation

type App struct {
	Processor
}

func NewApp(s Processor) *App {
	return &App{Processor: s}
}

func (a *App) RunApp() {
	a.Processor.Run()
}

type Processor interface {
	Run()
}
