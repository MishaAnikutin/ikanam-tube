package app

import "github.com/MishaAnikutin/metadata-api/internal/server"

type App struct {
	s server.Server
}

func New(server server.Server) *App {
	return &App{s: server}
}

func (app *App) Run() {
	app.s.Run()
}
