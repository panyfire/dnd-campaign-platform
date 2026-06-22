package app

import "github.com/panyfire/dnd-campaign-platform/internal/shared/config"

type App struct {
	Config *config.Config
}

func New(cfg *config.Config) *App {
	return &App{
		Config: cfg,
	}
}
