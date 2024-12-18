package app

import (
	"context"
	"log"

	"github.com/CaptainFallaway/XDH/internal"
	"github.com/CaptainFallaway/XDH/internal/grouping"
	"github.com/CaptainFallaway/XDH/internal/parsers"
	"github.com/CaptainFallaway/XDH/internal/storage"
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

	Current *storage.Session
}

func NewApp() *App {
	store := Must(storage.NewStorageService())

	return &App{
		Storage: store,
	}
}

// Wails specific context retrieval
func (app *App) OnStartup(ctx context.Context) {
	app.Ctx = ctx
}

// Must provides a centralized way of handling critical calls
// This way i can much easier create error messages. As an example
// If the storage does not want to instantiate, or we get some weird uuid error.
func Must[T any](val T, err error) T {
	if err != nil {
		log.Fatal(err)
	}
	return val
}

func (app *App) OpenFileDialog() string {
	return Must(runtime.OpenFileDialog(app.Ctx, dialogOptions))
}

// NewSession creates a empty [internal.Session] but with the uid set
// To a new uuidV7.
func (app *App) NewSessionData() internal.SessionData {
	uid := Must(uuid.NewV7())
	return internal.SessionData{
		Uid: uid.String(),
	}
}

func (app *App) NewSession(sessionData internal.SessionData, path string) {
	parsed := Must(parsers.Parse(path))

	grouped := grouping.MakeBoatGroupings(parsed)

	session := storage.NewStore(&sessionData, grouped)

	err := app.Storage.SetStore(session)
	// TODO: Handle error
	if err != nil {
		log.Fatal(err)
	}

	app.Current = session
}

func (app *App) SetSession(uid string) {
	session := Must(app.Storage.GetStore(uid))
	app.Current = session
}

func (app *App) ListSessions() []internal.SessionData {
	sessions, err := app.Storage.ListStores()
	if err != nil {
		log.Fatal(err)
	}

	return sessions
}

func (app *App) GetGroupings(sortingMetal string) []internal.Grouping {
	internal.SortByViolations(app.Current.Groupings, sortingMetal)
	return app.Current.Groupings
}
