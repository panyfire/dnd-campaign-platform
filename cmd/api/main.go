package main

import (
	"fmt"
	"os"

	"github.com/panyfire/dnd-campaign-platform/internal/app"
	"github.com/panyfire/dnd-campaign-platform/internal/shared/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Println("Could not load config:", err)
		os.Exit(1)
	}

	application := app.New(cfg)

	fmt.Println("Starting Chronicle...")
	fmt.Printf("Enviroment: %s\n", application.Config.AppEnv)
}
