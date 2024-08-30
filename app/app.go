package app

import (
	"context"
	"fmt"
	"strings"

	"github.com/CaptainFallaway/XDH/data_pipeline"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	Ctx       context.Context
	Groupings []data_pipeline.Grouping
}

func NewApp() *App {
	return &App{}
}

func (app *App) Startup(ctx context.Context) {
	app.Ctx = ctx
}

func (app *App) loadScans(path string) error {
	var (
		err   error
		scans []data_pipeline.ScanRow
	)

	if strings.HasSuffix(path, ".csv") {
		scans, err = data_pipeline.ParseCsv(path)
	} else if strings.HasSuffix(path, ".xlsx") || strings.HasSuffix(path, ".xls") {
		scans, err = data_pipeline.ParseExcel(path)
	} else {
		return fmt.Errorf("invalid file type")
	}

	app.Groupings = data_pipeline.GroupByBoat(&scans)

	return err
}

func (app *App) OpenFileDialog() {
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
		fmt.Printf("Error loading file: %s \n", err.Error())
		return
	}

	runtime.EventsEmit(app.Ctx, "modelLoaded")
}

func (app *App) GetModels(sortingMetal string) []data_pipeline.Grouping {
	data_pipeline.SortByViolations(&app.Groupings, sortingMetal)
	return app.Groupings
}
