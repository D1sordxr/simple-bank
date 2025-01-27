package presentation

// TODO: App container and Run() method

//func (a *App) Run(port string) error {
//
//	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
//
//	if err != nil {
//		log.Fatalf("Failed to listen: %v", err)
//	}
//	reflection.Register(s)
//	g.server = s
//	err = s.Serve(lis)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//func (a *App) Down() {
//	g.server.GracefulStop()
//}

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
