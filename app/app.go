package app

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/CaptainFallaway/XDH/data_pipeline"
	"github.com/CaptainFallaway/XDH/templates"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	Ctx          context.Context
	SortingMetal string
	Scans        []data_pipeline.ScanRow
	Loaded       bool
}

func NewApp() *App {
	return &App{
		SortingMetal: "Sn",
	}
}

func (app *App) Startup(ctx context.Context) {
	app.Ctx = ctx
}

func (app *App) loadScans(path string) error {
	var err error
	app.Loaded = false

	if strings.HasSuffix(path, ".csv") {
		app.Scans, err = data_pipeline.ParseCsv(path)
		app.Loaded = true
	} else if strings.HasSuffix(path, ".xlsx") || strings.HasSuffix(path, ".xls") {
		app.Scans, err = data_pipeline.ParseExcel(path)
		app.Loaded = true
	} else {
		return fmt.Errorf("invalid file type")
	}

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
}

func (app *App) GetBoatIDModel(index string, expanded string) string {
	convIndex, err := strconv.ParseUint(index, 10, 32)
	if err != nil {
		panic(err)
	}

	convExpanded, err := strconv.ParseBool(expanded)
	if err != nil {
		panic(err)
	}

	grouping := data_pipeline.GroupByBoat(&app.Scans)

	for _, group := range grouping {
		if group.Index == uint32(convIndex) {
			return Render(templates.BoatIDModelContent(group, !convExpanded))
		}
	}

	return "Something went horribly wrong!"
}

func (app *App) GroupByBoatID() string {
	grouping := data_pipeline.GroupByBoat(&app.Scans)

	data_pipeline.SortByViolations(&grouping, app.SortingMetal)

	return Render(templates.BoatIDModelList(&grouping, app.SortingMetal))
}

func (app *App) SetSortingMetal(metal string) string {
	if MetalInMetalPolicy(metal) {
		app.SortingMetal = metal
	}

	return app.GroupByBoatID()
}
