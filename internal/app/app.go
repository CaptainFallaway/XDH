package app

import (
	"context"
	"fmt"

	"github.com/CaptainFallaway/XDH/internal"
	"github.com/CaptainFallaway/XDH/internal/storage"
	"github.com/adrg/xdg"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	Ctx context.Context

	surveysPath   string
	groupingsPath string

	Storage storage.Storage
}

func NewApp() (*App, error) {
	surveysPath, err := xdg.DataFile(fmt.Sprintf("%s/sessions", internal.AppName))
	if err != nil {
		return nil, err
	}

	groupingsPath, err := xdg.DataFile(fmt.Sprintf("%s/groupings", internal.AppName))
	if err != nil {
		return nil, err
	}

	return &App{
		surveysPath:   surveysPath,
		groupingsPath: groupingsPath,
	}, nil
}

// Wails specific context retrieval
func (app *App) OnStartup(ctx context.Context) {
	app.Ctx = ctx

	storage := storage.NewTestStorage()

	app.Storage = storage
	runtime.LogInfo(app.Ctx, "Storage initialized")
}

func (app *App) OnShutdown(ctx context.Context) {
	runtime.LogDebug(app.Ctx, "Shutting down")

	err := app.Storage.Close()

	if err != nil {
		runtime.LogFatal(app.Ctx, err.Error())
	}
}

// HandleError TODO
func (app *App) HandleError(err error) {
	runtime.LogFatal(app.Ctx, fmt.Sprintf("Error Error Error: %s", err.Error()))
}

// OpenFileDialog opens a file dialog through the wails API, returning a path to the selected file.
func (app *App) OpenFileDialog() string {
	path, err := runtime.OpenFileDialog(app.Ctx, runtime.OpenDialogOptions{
		ShowHiddenFiles: true,
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Excel or Csv files",
				Pattern:     "*.xlsx;*.xls;*.csv",
			},
		},
	})

	if err != nil {
		runtime.LogFatal(app.Ctx, err.Error())
		return ""
	}

	return path
}
