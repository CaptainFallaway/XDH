package app

import (
	"github.com/CaptainFallaway/XDH/internal"
	"github.com/CaptainFallaway/XDH/internal/api"
	"github.com/CaptainFallaway/XDH/internal/models"
)

// createGroupings inserts the groupings into the storage
// and returns an error if it fails.
//
// This method is used when creating a new survey
// and when updating an existing survey.
func (app *App) insertGroupings(groupings []models.Grouping) error {
	for _, grouping := range groupings {
		err := app.Storage.InsertGrouping(&grouping)
		if err != nil {
			return err
		}
	}
	return nil
}

func (app *App) UpdateGrouping(groupingId string, grouping *api.Grouping) {
	err := app.Storage.UpdateGrouping(groupingId, grouping)
	if err != nil {
		app.HandleError(err)
		return
	}
}

func (app *App) GetGroupings(groupingIds []string, sortingMetal string) []models.Grouping {
	groupings, err := app.Storage.GetGroupings(groupingIds)
	if err != nil {
		app.HandleError(err)
		return nil
	}

	// Sort the groupings by violations
	internal.SortByViolations(groupings, sortingMetal)

	return groupings
}
