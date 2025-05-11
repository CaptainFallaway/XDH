package app

import (
	"context"
	"fmt"
	"sync"

	"github.com/CaptainFallaway/XDH/internal"
	"github.com/CaptainFallaway/XDH/internal/storage"
	"github.com/adrg/xdg"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	Mux sync.Mutex

	Ctx context.Context

	Sessions  storage.KeyValueStorage[*internal.Session]
	Groupings storage.KeyValueStorage[[]internal.Grouping]
}

func NewApp() (*App, error) {
	sessionsPath, err := xdg.DataFile(fmt.Sprintf("%s/sessions", internal.AppName))
	if err != nil {
		return nil, err
	}

	sessions, err := storage.NewBadgerStorage[*internal.Session](sessionsPath)
	if err != nil {
		return nil, err
	}

	groupingsPath, err := xdg.DataFile(fmt.Sprintf("%s/groupings", internal.AppName))
	if err != nil {
		return nil, err
	}

	groupings, err := storage.NewBadgerStorage[[]internal.Grouping](groupingsPath)
	if err != nil {
		return nil, err
	}

	return &App{
		Sessions:  sessions,
		Groupings: groupings,
	}, nil
}

// Wails specific context retrieval
func (app *App) OnStartup(ctx context.Context) {
	app.Ctx = ctx
}

func (app *App) OnShutdown(ctx context.Context) {
	runtime.LogDebug(app.Ctx, "Shutting down")

	err := app.Sessions.Close()

	if err != nil {
		runtime.LogFatal(app.Ctx, err.Error())
	}

	err = app.Groupings.Close()

	if err != nil {
		runtime.LogFatal(app.Ctx, err.Error())
	}
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
