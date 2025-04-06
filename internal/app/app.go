package app

import (
	"context"
	"os"

	"github.com/CaptainFallaway/XDH/internal"
	"github.com/CaptainFallaway/XDH/internal/grouping"
	"github.com/CaptainFallaway/XDH/internal/parsers"
	"github.com/CaptainFallaway/XDH/internal/storage"
	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"github.com/wailsapp/wails/v2/pkg/runtime"

	"sync"
)

var dialogOptions = runtime.OpenDialogOptions{
	ShowHiddenFiles: true,
	Filters: []runtime.FileFilter{
		{
			DisplayName: "Excel or Csv files",
			Pattern:     "*.xlsx;*.xls;*.csv",
		},
	},
}

type App struct {
	Mux sync.Mutex
	Ctx context.Context

	Storage storage.StorageService

	Log *log.Logger
}

func NewApp() *App {
	return &App{}
}

// Wails specific context retrieval
func (app *App) OnStartup(ctx context.Context) {
	app.Ctx = ctx
	app.Storage, _ = storage.NewStorageService() // TODO: Handle error
	app.Log = log.NewWithOptions(os.Stderr, log.Options{
		ReportTimestamp: true,
		ReportCaller:    true,
	})
}

func (app *App) OnShutdown(ctx context.Context) {
	app.Log.Debug("Shutting down")
	err := app.Storage.Close()
	if err != nil {
		app.Log.Fatal(err)
	}
}

// Must provides a centralized way of handling critical calls
// This way i can much easier create error messages. As an example
// If the storage does not want to instantiate, or we get some weird uuid error.
// func Must[T any](val T, err error) T {
// 	if err != nil {
// 		app.Log.Fatal(err)
// 	}
// 	return val
// }

// OpenFileDiaapp.Log opens a file dialog with [dialogoptions] and returns the path to the file
func (app *App) OpenFileDialog() string {
	path, err := runtime.OpenFileDialog(app.Ctx, dialogOptions)
	if err != nil {
		app.Log.Fatal(err)
		return ""
	}
	return path
}

// NewSessionData creates an empty [internal.SessionInfo] but with the uid set
// To a new uuidV7.
func (app *App) NewSessionData() internal.SessionInfo {
	uid, err := uuid.NewV7()
	if err != nil {
		app.Log.Fatal(err)
		return internal.SessionInfo{}
	}
	return internal.SessionInfo{
		Uid: uid.String(),
	}
}

func (app *App) CreateSession(sessionData internal.SessionInfo, path string) string {
	parsed, err := parsers.Parse(path)

	if err != nil {
		app.Log.Fatal(err)
		return ""
	}

	grouped := grouping.MakeBoatGroupings(parsed)

	session := internal.NewSession(&sessionData, grouped)

	err = app.Storage.Set(session)
	if err != nil {
		app.Log.Fatal(err)
		return ""
	}

	return sessionData.Uid
}

func (app *App) DeleteSession(uid string) {
	err := app.Storage.Delete(uid)
	if err != nil {
		app.Log.Fatal(err)
	}
}

func (app *App) ListSessions() []internal.SessionInfo {
	sessions, err := app.Storage.List()

	if err != nil {
		app.Log.Fatal(err)
		return nil
	}

	return sessions
}

// func (app *App) SetSession(uid string) {
// 	session := Must(app.Storage.Get(uid))
// 	app.Current = session
// }

func (app *App) GetGroupings(sessionId string, sortingMetal string) []internal.Grouping {
	if sortingMetal == "" || sessionId == "" {
		return nil
	}

	session, err := app.Storage.Get(sessionId)

	if err != nil {
		app.Log.Fatal(err)
		return nil
	}

	internal.SortByViolations(session.Groupings, sortingMetal)
	return session.Groupings
}

func (app *App) GetSessionData(sessionId string) internal.SessionInfo {
	session, err := app.Storage.Get(sessionId)
	if err != nil {
		app.Log.Fatal(err)
		return internal.SessionInfo{}
	}
	return *session.Session
}

func (app *App) SetSessionData(sessionData internal.SessionInfo) {
	session, err := app.Storage.Get(sessionData.Uid)

	if err != nil {
		app.Log.Fatal(err)
		return
	}

	session.Session = &sessionData

	err = app.Storage.Set(session)
	if err != nil {
		app.Log.Fatal(err)
	}
}

// func (app *App) UpdateGrouping(grouping internal.Grouping) {
// 	if app.Current == nil {
// 		app.Log.Fatal("No session selected")
// 	}

// 	groupings := app.Current.Groupings

// 	for i, g := range groupings {
// 		if g.BoatID == grouping.BoatID {
// 			groupings[i] = grouping
// 		}
// 	}

// 	app.setToDatabase()
// }

// setToDatabase just sets the `Current` field to the database
// func (app *App) setToDatabase() {
// 	err := app.Storage.Set(app.Current)
// 	if err != nil {
// 		app.Log.Fatal(err)
// 	}
// }
