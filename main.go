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
	application := app.NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:     "XRF",
		Width:     1024,
		Height:    768,
		MinWidth:  606,
		MinHeight: 400,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: application.Startup,
		Bind: []interface{}{
			application,
			new(app.ModelInterface),
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
