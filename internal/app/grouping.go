package app

import (
	"github.com/CaptainFallaway/XDH/internal"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// GetGroupings returns a list of all groupings from a session
func (app *App) GetGroupings(uid string, sortingMetal string) []internal.Grouping {
	if sortingMetal == "" || uid == "" {
		return nil
	}

	groupings, err := app.Groupings.Get(uid)
	if err != nil {
		runtime.LogFatal(app.Ctx, err.Error())
		return nil
	}

	internal.SortByViolations(groupings, sortingMetal)
	return groupings
}

// UpdateGrouping updates the singular grouping in the full collection of groupings
func (app *App) UpdateGrouping(uid string, grouping internal.Grouping) {}
