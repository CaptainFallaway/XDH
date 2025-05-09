package app

import (
	"context"
	"fmt"
	"sync"

	"github.com/CaptainFallaway/XDH/internal"
	"github.com/CaptainFallaway/XDH/internal/grouping"
	"github.com/CaptainFallaway/XDH/internal/parsers"
	"github.com/CaptainFallaway/XDH/internal/storage"
	"github.com/adrg/xdg"
	"github.com/google/uuid"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	Mux sync.Mutex
	Ctx context.Context

	Storage storage.StorageService
}

func NewApp() (*App, error) {
	path, err := xdg.DataFile(internal.AppName)
	if err != nil {
		return nil, err
	}

	fmt.Println(path)

	store, err := storage.NewBadgerStorage(path) // TODO: Handle error
	if err != nil {
		return nil, err
	}

	return &App{Storage: store}, nil
}

// Wails specific context retrieval
func (app *App) OnStartup(ctx context.Context) {
	app.Ctx = ctx
}

func (app *App) OnShutdown(ctx context.Context) {
	runtime.LogDebug(app.Ctx, "Shutting down")

	err := app.Storage.Close()

	if err != nil {
		runtime.LogFatal(app.Ctx, err.Error())
	}
}

// OpenFileDiaapp.Log opens a file dialog with [dialogoptions] and returns the path to the file
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

// NewSessionData creates an empty [internal.SessionInfo] but with the uid set
// To a new uuidV7.
func (app *App) NewSessionData() internal.SessionInfo {
	uid, err := uuid.NewV7()
	if err != nil {
		runtime.LogFatal(app.Ctx, err.Error())
		return internal.SessionInfo{}
	}
	return internal.SessionInfo{
		Uid: uid.String(),
	}
}

func (app *App) CreateSession(sessionData internal.SessionInfo, path string) string {
	parsed, err := parsers.Parse(path)

	if err != nil {
		runtime.LogFatal(app.Ctx, err.Error())
		return ""
	}

	grouped := grouping.MakeBoatGroupings(parsed)

	session := internal.NewSession(&sessionData, grouped)

	err = app.Storage.Set(session)
	if err != nil {
		runtime.LogFatal(app.Ctx, err.Error())
		return ""
	}

	return sessionData.Uid
}

func (app *App) DeleteSession(uid string) {
	err := app.Storage.Delete(uid)
	if err != nil {
		runtime.LogFatal(app.Ctx, err.Error())
	}
}

func (app *App) ListSessions() []internal.SessionInfo {
	sessions, err := app.Storage.List()

	if err != nil {
		runtime.LogFatal(app.Ctx, err.Error())
		return nil
	}

	return sessions
}

func (app *App) GetGroupings(sessionId string, sortingMetal string) []internal.Grouping {
	if sortingMetal == "" || sessionId == "" {
		return nil
	}

	session, err := app.Storage.Get(sessionId)
	if err != nil {
		runtime.LogFatal(app.Ctx, err.Error())
		return nil
	}

	internal.SortByViolations(session.Groupings, sortingMetal)
	return session.Groupings
}

func (app *App) GetSessionData(sessionId string) internal.SessionInfo {
	session, err := app.Storage.Get(sessionId)
	if err != nil {
		runtime.LogFatal(app.Ctx, err.Error())
		return internal.SessionInfo{}
	}

	return *session.Session
}

func (app *App) SetSessionData(sessionData internal.SessionInfo) {
	session, err := app.Storage.Get(sessionData.Uid)
	if err != nil {
		runtime.LogFatal(app.Ctx, err.Error())
		return
	}

	session.Session = &sessionData

	err = app.Storage.Set(session)
	if err != nil {
		runtime.LogFatal(app.Ctx, err.Error())
	}
}
