package app

import (
	"context"
	"log"

	"github.com/CaptainFallaway/XDH/internal"
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

	SessionStorage storage.StorageService

	Groupings []internal.Grouping
	Session   internal.Session
}

func NewApp(dataDir string) *App {
	store := Must(storage.NewStorageService(dataDir))

	return &App{
		SessionStorage: store,
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

// NewSession creates a empty [internal.Session] but with the uid set
// To a new uuidV7.
func (app *App) NewSession() internal.Session {
	uid := Must(uuid.NewV7())
	return internal.Session{
		Uid: uid.String(),
	}
}

func (app *App) OpenFileDialog() string {
	return Must(runtime.OpenFileDialog(app.Ctx, dialogOptions))
}

func (app *App) InitializeSession(session internal.Session, path string) []internal.Grouping {
	return nil
}
