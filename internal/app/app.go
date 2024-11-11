package app

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/CaptainFallaway/XDH/internal"
	"github.com/CaptainFallaway/XDH/internal/grouping"
	"github.com/CaptainFallaway/XDH/internal/parsers"

	"github.com/wailsapp/wails/v2/pkg/runtime"

	"sync"
)

type App struct {
	Mux       sync.Mutex
	Ctx       context.Context
	Groupings []internal.Grouping
}

func NewApp() *App {
	return &App{}
}

// OnStartup Retrieves the wails runtime context for wails runtime methods
func (app *App) OnStartup(ctx context.Context) {
	app.Ctx = ctx
}

func (app *App) loadScans(path string) error {
	var (
		err   error
		scans *[]internal.ScanRow
	)

	if strings.HasSuffix(path, ".csv") {
		scans, err = parsers.ParseCsvFile(path)
	} else if strings.HasSuffix(path, ".xlsx") || strings.HasSuffix(path, ".xls") {
		scans, err = parsers.ParseExcelFile(path)
	} else {
		return fmt.Errorf("invalid file type")
	}

	if err != nil {
		log.Println(err)
		panic("here is johnny")
	}

	app.Groupings = grouping.MakeBoatGroupings(scans)

	return err
}

func (app *App) OpenFileDialog() {
	app.Mux.Lock()
	defer app.Mux.Unlock()

	path, err := runtime.OpenFileDialog(app.Ctx, dialogOptions)

	if err != nil {
		// Make some event thing for toasts on the frontend to display errors
		fmt.Printf("Error opening file: %s \n", err.Error())
		return
	}

	if path == "" {
		return
	}

	err = app.loadScans(path)

	if err != nil {
		// Make some event thing for toasts on the frontend to display errors
		log.Printf("Error loading file: %s \n", err)
		return
	}
}

func (app *App) GetModels(sortingMetal string) []internal.Grouping {
	app.Mux.Lock()
	defer app.Mux.Unlock()

	internal.SortByViolations(&app.Groupings, sortingMetal)
	return app.Groupings
}
