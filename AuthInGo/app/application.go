package app



type Config struct {
	Addr string
}

type Application struct {
	Config
}

func (app *Application) Run() error {
	server :=
}