package main

import (
	"embed"

	"github.com/CaptainFallaway/XDH/app"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure

	// We might want to just develop in the app.go file

	app := app.NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:     "XRF",
		Width:     1024,
		Height:    768,
		MinWidth:  700,
		MinHeight: 500,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: app.Startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
