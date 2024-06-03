package app

import (
	"context"
	"fmt"
	"strings"

	"github.com/CaptainFallaway/XDH/data_pipeline"
	"github.com/CaptainFallaway/XDH/templates"

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

	fmt.Printf("Path: %s", path)

	if path == "" {
		return
	}

	err = app.loadScans(path)

	if err != nil {
		// Make some event thing for toasts on the frontend to display errors
		fmt.Printf("Error loading file: %s \n", err.Error())
		return
	}

	fmt.Print("File loaded\n")
	runtime.EventsEmit(app.Ctx, "modelLoaded")
}

func (app *App) GetModels(sortingMetal string) string {
	data_pipeline.SortByViolations(&app.Groupings, sortingMetal)

	return Render(templates.BoatIDModelList(&app.Groupings, sortingMetal))
}

func (app *App) GetModelContent(index int) string {
	return ""
}

func (app *App) GetDropArea() string {
	return Render(templates.DropArea())
	// asd
}
