package main

import (
	"embed"
	"fmt"
	"os"

	"github.com/CaptainFallaway/XDH/internal"
	"github.com/CaptainFallaway/XDH/internal/app"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app, err := app.NewApp()
	if err != nil {
		fmt.Println("Error instantiating app: ", err.Error())
		os.Exit(1)
	}

	err = wails.Run(&options.App{
		Title:     internal.AppName,
		Width:     1024,
		Height:    768,
		MinWidth:  0,
		MinHeight: 0,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup:  app.OnStartup,
		OnShutdown: app.OnShutdown,
		Bind: []any{
			app,
		},
	})

	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
}
