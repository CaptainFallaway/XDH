package main

import (
	"embed"

	"github.com/CaptainFallaway/XDH/internal"
	"github.com/CaptainFallaway/XDH/internal/app"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app := app.NewApp()

	err := wails.Run(&options.App{
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
		println("Error:", err.Error())
	}
}
