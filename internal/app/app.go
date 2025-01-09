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

	Current *internal.Session
}

func NewApp() *App {
	return &App{}
}

// Wails specific context retrieval
func (app *App) OnStartup(ctx context.Context) {
	app.Ctx = ctx
	app.Storage = Must(storage.NewStorageService()) // TODO: Add dev mode
}

func (app *App) OnShutdown(ctx context.Context) {
	log.Println("Shutting down")
	err := app.Storage.Close()
	if err != nil {
		log.Println(err)
	}
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

// OpenFileDialog opens a file dialog with [dialogoptions] and returns the path to the file
func (app *App) OpenFileDialog() string {
	return Must(runtime.OpenFileDialog(app.Ctx, dialogOptions))
}

// NewSessionData creates an empty [internal.SessionInfo] but with the uid set
// To a new uuidV7.
func (app *App) NewSessionData() internal.SessionInfo {
	uid := Must(uuid.NewV7())
	return internal.SessionInfo{
		Uid: uid.String(),
	}
}

func (app *App) CreateSession(sessionData internal.SessionInfo, path string) {
	parsed := Must(parsers.Parse(path))

	grouped := grouping.MakeBoatGroupings(parsed)

	session := internal.NewSession(&sessionData, grouped)

	app.Current = session

	err := app.Storage.Set(session)
	if err != nil {
		log.Fatal(err)
	}
}

func (app *App) DeleteSession(uid string) {
	err := app.Storage.Delete(uid)
	if err != nil {
		log.Fatal(err)
	}
}

func (app *App) ListSessions() []internal.SessionInfo {
	return Must(app.Storage.List())
}

func (app *App) SetSession(uid string) {
	session := Must(app.Storage.Get(uid))
	app.Current = session
}

func (app *App) GetGroupings(sortingMetal string) []internal.Grouping {
	if app.Current == nil {
		log.Fatal("No session selected")
	}

	internal.SortByViolations(app.Current.Groupings, sortingMetal)
	return app.Current.Groupings
}

// setToDatabase just sets the `Current` field to the database
func (app *App) setToDatabase() {
	err := app.Storage.Set(app.Current)
	if err != nil {
		log.Fatal(err)
	}
}

func (app *App) UpdateSessionData(sessionData internal.SessionInfo) {
	if app.Current == nil {
		log.Fatal("No session selected")
	}

	app.Current.Session = &sessionData
	app.setToDatabase()
}

func (app *App) UpdateGrouping(grouping internal.Grouping) {
	if app.Current == nil {
		log.Fatal("No session selected")
	}

	groupings := app.Current.Groupings

	for i, g := range groupings {
		if g.BoatID == grouping.BoatID {
			groupings[i] = grouping
		}
	}

	app.setToDatabase()
}
